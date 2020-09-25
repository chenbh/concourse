package atc

import "github.com/concourse/concourse/atc/types"

type BuildPreparation struct {
	BuildID             int                                     `json:"build_id"`
	PausedPipeline      types.BuildPreparationStatus            `json:"paused_pipeline"`
	PausedJob           types.BuildPreparationStatus            `json:"paused_job"`
	MaxRunningBuilds    types.BuildPreparationStatus            `json:"max_running_builds"`
	Inputs              map[string]types.BuildPreparationStatus `json:"inputs"`
	InputsSatisfied     types.BuildPreparationStatus            `json:"inputs_satisfied"`
	MissingInputReasons types.MissingInputReasons               `json:"missing_input_reasons"`
}
