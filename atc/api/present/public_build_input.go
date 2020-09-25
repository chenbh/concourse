package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func PublicBuildInput(input db.BuildInput, pipelineID int) types.PublicBuildInput {
	return types.PublicBuildInput{
		Name:            input.Name,
		Version:         types.Version(input.Version),
		PipelineID:      pipelineID,
		FirstOccurrence: input.FirstOccurrence,
	}
}

func PublicBuildOutput(output db.BuildOutput) types.PublicBuildOutput {
	return types.PublicBuildOutput{
		Name:    output.Name,
		Version: types.Version(output.Version),
	}
}
