package configserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"io/ioutil"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/configvalidate"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/vars"
	"github.com/hashicorp/go-multierror"
	"github.com/tedsuo/rata"
)

func (s *Server) SaveConfig(w http.ResponseWriter, r *http.Request) {
	session := s.logger.Session("set-config")

	query := r.URL.Query()

	checkCredentials := false
	if _, exists := query[types.SaveConfigCheckCreds]; exists {
		checkCredentials = true
	}

	var version db.ConfigVersion
	if configVersionStr := r.Header.Get(types.ConfigVersionHeader); len(configVersionStr) != 0 {
		_, err := fmt.Sscanf(configVersionStr, "%d", &version)
		if err != nil {
			session.Error("malformed-config-version", err)
			s.handleBadRequest(w, fmt.Sprintf("config version is malformed: %s", err))
			return
		}
	}

	var config types.Config
	switch r.Header.Get("Content-type") {
	case "application/json", "application/x-yaml":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.handleBadRequest(w, fmt.Sprintf("read failed: %s", err))
			return
		}

		err = types.UnmarshalConfig(body, &config)
		if err != nil {
			session.Error("malformed-request-payload", err, lager.Data{
				"content-type": r.Header.Get("Content-Type"),
			})

			s.handleBadRequest(w, fmt.Sprintf("malformed config: %s", err))
			return
		}
	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	warnings, errorMessages := configvalidate.Validate(config)
	if len(errorMessages) > 0 {
		session.Info("ignoring-invalid-config", lager.Data{"errors": errorMessages})
		s.handleBadRequest(w, errorMessages...)
		return
	}

	pipelineName := rata.Param(r, "pipeline_name")
	warning := types.ValidateIdentifier(pipelineName, "pipeline")
	if warning != nil {
		warnings = append(warnings, *warning)
	}

	teamName := rata.Param(r, "team_name")
	warning = types.ValidateIdentifier(teamName, "team")
	if warning != nil {
		warnings = append(warnings, *warning)
	}

	if checkCredentials {
		variables := creds.NewVariables(s.secretManager, teamName, pipelineName, false)

		errs := validateCredParams(variables, config, session)
		if errs != nil {
			s.handleBadRequest(w, fmt.Sprintf("credential validation failed\n\n%s", errs))
			return
		}
	}

	session.Info("saving")

	team, found, err := s.teamFactory.FindTeam(teamName)
	if err != nil {
		session.Error("failed-to-find-team", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !found {
		session.Debug("team-not-found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, created, err := team.SavePipeline(pipelineName, config, version, true)
	if err != nil {
		session.Error("failed-to-save-config", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to save config: %s", err)
		return
	}

	if !created {
		if err = s.teamFactory.NotifyResourceScanner(); err != nil {
			session.Error("failed-to-notify-resource-scanner", err)
		}
	}

	session.Info("saved")

	w.Header().Set("Content-Type", "application/json")

	if created {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	s.writeSaveConfigResponse(w, types.SaveConfigResponse{Warnings: warnings})
}

// Simply validate that the credentials exist; don't do anything with the actual secrets
func validateCredParams(credMgrVars vars.Variables, config types.Config, session lager.Logger) error {
	var errs error

	for _, resourceType := range config.ResourceTypes {
		_, err := creds.NewSource(credMgrVars, resourceType.Source).Evaluate()
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	for _, resource := range config.Resources {
		_, err := creds.NewSource(credMgrVars, resource.Source).Evaluate()
		if err != nil {
			errs = multierror.Append(errs, err)
		}

		_, err = creds.NewString(credMgrVars, resource.WebhookToken).Evaluate()
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	for _, job := range config.Jobs {
		_ = job.StepConfig().Visit(types.StepRecursor{
			OnTask: func(step *types.TaskStep) error {
				err := creds.NewTaskEnvValidator(credMgrVars, step.Params).Validate()
				if err != nil {
					errs = multierror.Append(errs, err)
				}

				err = creds.NewTaskVarsValidator(credMgrVars, step.Vars).Validate()
				if err != nil {
					errs = multierror.Append(errs, err)
				}

				if step.Config != nil {
					// embedded task - we can fully validate it, interpolating with cred mgr variables
					var taskConfigSource exec.TaskConfigSource
					embeddedTaskVars := []vars.Variables{credMgrVars}
					taskConfigSource = exec.StaticConfigSource{Config: step.Config}
					taskConfigSource = exec.InterpolateTemplateConfigSource{
						ConfigSource:  taskConfigSource,
						Vars:          embeddedTaskVars,
						ExpectAllKeys: true,
					}
					taskConfigSource = exec.ValidatingConfigSource{ConfigSource: taskConfigSource}
					_, err = taskConfigSource.FetchConfig(context.TODO(), session, nil)
					if err != nil {
						errs = multierror.Append(errs, err)
					}
				}

				return nil
			},
		})
	}

	if errs != nil {
		session.Info("config-has-invalid-creds", lager.Data{"errors": errs.Error()})
	}

	return errs
}

func (s *Server) handleBadRequest(w http.ResponseWriter, errorMessages ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	s.writeSaveConfigResponse(w, types.SaveConfigResponse{
		Errors: errorMessages,
	})
}

func (s *Server) writeSaveConfigResponse(w http.ResponseWriter, saveConfigResponse types.SaveConfigResponse) {
	responseJSON, err := json.Marshal(saveConfigResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to generate error response: %s", err)
		return
	}

	_, err = w.Write(responseJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
