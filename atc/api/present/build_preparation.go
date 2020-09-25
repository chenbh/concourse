package present

import (
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func BuildPreparation(preparation db.BuildPreparation) atc.BuildPreparation {
	inputs := make(map[string]types.BuildPreparationStatus)

	for k, v := range preparation.Inputs {
		inputs[k] = types.BuildPreparationStatus(v)
	}

	return atc.BuildPreparation{
		BuildID:             preparation.BuildID,
		PausedPipeline:      types.BuildPreparationStatus(preparation.PausedPipeline),
		PausedJob:           types.BuildPreparationStatus(preparation.PausedJob),
		MaxRunningBuilds:    types.BuildPreparationStatus(preparation.MaxRunningBuilds),
		Inputs:              inputs,
		InputsSatisfied:     types.BuildPreparationStatus(preparation.InputsSatisfied),
		MissingInputReasons: types.MissingInputReasons(preparation.MissingInputReasons),
	}
}
