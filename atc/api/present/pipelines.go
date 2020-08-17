package present

import (
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/db"
)

func Pipelines(savedPipelines []db.Pipeline) []atc.Pipeline {
	pipelines := make([]atc.Pipeline, len(savedPipelines))

	for i := range savedPipelines {
		pipelines[i] = Pipeline(savedPipelines[i])
	}

	return pipelines
}
