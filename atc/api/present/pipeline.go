package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Pipeline(savedPipeline db.Pipeline) types.Pipeline {
	return types.Pipeline{
		ID:          savedPipeline.ID(),
		Name:        savedPipeline.Name(),
		TeamName:    savedPipeline.TeamName(),
		Paused:      savedPipeline.Paused(),
		Public:      savedPipeline.Public(),
		Archived:    savedPipeline.Archived(),
		Groups:      savedPipeline.Groups(),
		Display:     savedPipeline.Display(),
		LastUpdated: savedPipeline.LastUpdated().Unix(),
	}
}
