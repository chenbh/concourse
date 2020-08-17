package present

import (
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/db"
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
