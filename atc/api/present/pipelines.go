package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Pipelines(savedPipelines []db.Pipeline) []types.Pipeline {
	pipelines := make([]types.Pipeline, len(savedPipelines))

	for i := range savedPipelines {
		pipelines[i] = Pipeline(savedPipelines[i])
	}

	return pipelines
}
