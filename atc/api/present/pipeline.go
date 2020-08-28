package present

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
)

func Pipeline(savedPipeline db.Pipeline) atc.Pipeline {
	return atc.Pipeline{
		ID:          savedPipeline.ID(),
		Name:        savedPipeline.Name(),
		TeamName:    savedPipeline.TeamName(),
		Paused:      savedPipeline.Paused(),
		Public:      savedPipeline.Public(),
		Archived:    savedPipeline.Archived(),
		Groups:      savedPipeline.Groups(),
		LastUpdated: savedPipeline.LastUpdated().Unix(),
	}
}
