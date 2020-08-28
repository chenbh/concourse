package present

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
)

func Pipelines(savedPipelines []db.Pipeline) []atc.Pipeline {
	pipelines := make([]atc.Pipeline, len(savedPipelines))

	for i := range savedPipelines {
		pipelines[i] = Pipeline(savedPipelines[i])
	}

	return pipelines
}
