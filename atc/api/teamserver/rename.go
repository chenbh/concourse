package teamserver

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"io/ioutil"
	"net/http"

	"github.com/concourse/concourse/atc/api/accessor"
)

// RenameTeam allows an authenticated user with authority or admin to rename a team
func (s *Server) RenameTeam(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("rename-team")
	acc := accessor.GetAccessor(r)

	teamName := r.FormValue(":team_name")
	if !acc.IsAdmin() && !acc.IsAuthorized(teamName) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("failed-to-read-body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var rename types.RenameRequest
	err = json.Unmarshal(data, &rename)
	if err != nil {
		logger.Error("failed-to-unmarshal-body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var warnings []types.ConfigWarning
	warning := types.ValidateIdentifier(rename.NewName, "team")
	if warning != nil {
		warnings = append(warnings, *warning)
	}

	team, found, err := s.teamFactory.FindTeam(teamName)
	if err != nil {
		logger.Error("failed-to-get-team", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !found {
		logger.Info("team-not-found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = team.Rename(rename.NewName)
	if err != nil {
		logger.Error("failed-to-update-team-name", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(types.SaveConfigResponse{Warnings: warnings})
	if err != nil {
		s.logger.Error("failed-to-encode-response", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
