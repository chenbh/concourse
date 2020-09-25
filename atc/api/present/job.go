package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Job(
	teamName string,
	job db.Job,
	inputs []types.JobInput,
	outputs []types.JobOutput,
	finishedBuild db.Build,
	nextBuild db.Build,
	transitionBuild db.Build,
) types.Job {
	var presentedNextBuild, presentedFinishedBuild, presentedTransitionBuild *types.Build

	if nextBuild != nil {
		presented := Build(nextBuild)
		presentedNextBuild = &presented
	}

	if finishedBuild != nil {
		presented := Build(finishedBuild)
		presentedFinishedBuild = &presented
	}

	if transitionBuild != nil {
		presented := Build(transitionBuild)
		presentedTransitionBuild = &presented
	}

	sanitizedInputs := []types.JobInput{}
	for _, input := range inputs {
		sanitizedInputs = append(sanitizedInputs, types.JobInput{
			Name:     input.Name,
			Resource: input.Resource,
			Passed:   input.Passed,
			Trigger:  input.Trigger,
		})
	}

	sanitizedOutputs := []types.JobOutput{}
	for _, output := range outputs {
		sanitizedOutputs = append(sanitizedOutputs, types.JobOutput{
			Name:     output.Name,
			Resource: output.Resource,
		})
	}

	return types.Job{
		ID: job.ID(),

		Name:                 job.Name(),
		PipelineName:         job.PipelineName(),
		TeamName:             teamName,
		DisableManualTrigger: job.DisableManualTrigger(),
		Paused:               job.Paused(),
		FirstLoggedBuildID:   job.FirstLoggedBuildID(),
		FinishedBuild:        presentedFinishedBuild,
		NextBuild:            presentedNextBuild,
		TransitionBuild:      presentedTransitionBuild,
		HasNewInputs:         job.HasNewInputs(),

		Inputs:  sanitizedInputs,
		Outputs: sanitizedOutputs,

		Groups: job.Tags(),
	}
}
