package setpipelinehelpers

import (
	"fmt"
	"net/url"
	"os"
	"sigs.k8s.io/yaml"

	"github.com/vito/go-interact/interact"

	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/configvalidate"
	"github.com/chenbh/concourse/v6/fly/commands/internal/displayhelpers"
	"github.com/chenbh/concourse/v6/fly/commands/internal/templatehelpers"
	"github.com/chenbh/concourse/v6/fly/rc"
	"github.com/chenbh/concourse/v6/fly/ui"
	"github.com/chenbh/concourse/v6/go-concourse/concourse"
)

type ATCConfig struct {
	PipelineName     string
	Team             concourse.Team
	TargetName       rc.TargetName
	Target           string
	SkipInteraction  bool
	CheckCredentials bool
	CommandWarnings  []concourse.ConfigWarning
}

func (atcConfig ATCConfig) ApplyConfigInteraction() bool {
	if atcConfig.SkipInteraction {
		return true
	}

	confirm := false
	err := interact.NewInteraction("apply configuration?").Resolve(&confirm)
	if err != nil {
		return false
	}

	return confirm
}

func (atcConfig ATCConfig) Set(yamlTemplateWithParams templatehelpers.YamlTemplateWithParams) error {
	evaluatedTemplate, err := yamlTemplateWithParams.Evaluate(false, false)
	if err != nil {
		return err
	}

	existingConfig, existingConfigVersion, _, err := atcConfig.Team.PipelineConfig(atcConfig.PipelineName)
	if err != nil {
		return err
	}

	var newConfig atc.Config
	err = yaml.Unmarshal([]byte(evaluatedTemplate), &newConfig)
	if err != nil {
		return err
	}

	configWarnings, _ := configvalidate.Validate(newConfig)
	for _, w := range configWarnings {
		atcConfig.CommandWarnings = append(atcConfig.CommandWarnings, concourse.ConfigWarning{
			Type:    w.Type,
			Message: w.Message,
		})
	}

	diffExists := diff(existingConfig, newConfig)

	if len(atcConfig.CommandWarnings) > 0 {
		displayhelpers.ShowWarnings(atcConfig.CommandWarnings)
	}

	if !diffExists {
		fmt.Println("no changes to apply")
		return nil
	}

	if !atcConfig.ApplyConfigInteraction() {
		fmt.Println("bailing out")
		return nil
	}

	created, updated, warnings, err := atcConfig.Team.CreateOrUpdatePipelineConfig(
		atcConfig.PipelineName,
		existingConfigVersion,
		evaluatedTemplate,
		atcConfig.CheckCredentials,
	)
	if err != nil {
		return err
	}

	updatedConfig, found, err := atcConfig.Team.Pipeline(atcConfig.PipelineName)
	if err != nil {
		return err
	}

	paused := found && updatedConfig.Paused

	if len(warnings) > 0 {
		displayhelpers.ShowWarnings(warnings)
	}

	atcConfig.showPipelineUpdateResult(created, updated, paused)
	return nil
}

func (atcConfig ATCConfig) UnpausePipelineCommand() string {
	return fmt.Sprintf("%s -t %s unpause-pipeline -p %s", os.Args[0], atcConfig.TargetName, atcConfig.PipelineName)
}

func (atcConfig ATCConfig) showPipelineUpdateResult(created bool, updated bool, paused bool) {
	if updated {
		fmt.Println("configuration updated")
	} else if created {

		targetURL, err := url.Parse(atcConfig.Target)
		if err != nil {
			fmt.Println("Could not parse targetURL")
		}

		pipelineURL, err := url.Parse("/teams/" + atcConfig.Team.Name() + "/pipelines/" + atcConfig.PipelineName)
		if err != nil {
			fmt.Println("Could not parse pipelineURL")
		}

		fmt.Println("pipeline created!")
		fmt.Printf("you can view your pipeline here: %s\n", targetURL.ResolveReference(pipelineURL))
	} else {
		panic("Something really went wrong!")
	}

	if paused {
		fmt.Println("")
		fmt.Println("the pipeline is currently paused. to unpause, either:")
		fmt.Println("  - run the unpause-pipeline command:")
		fmt.Println("    " + atcConfig.UnpausePipelineCommand())
		fmt.Println("  - click play next to the pipeline in the web ui")
	}
}

func diff(existingConfig atc.Config, newConfig atc.Config) bool {
	stdout, _ := ui.ForTTY(os.Stdout)
	return existingConfig.Diff(stdout, newConfig)
}
