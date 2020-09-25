package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func WorkerArtifacts(artifacts []db.WorkerArtifact) []types.WorkerArtifact {
	wa := []types.WorkerArtifact{}
	for _, a := range artifacts {
		wa = append(wa, WorkerArtifact(a))
	}
	return wa
}

func WorkerArtifact(artifact db.WorkerArtifact) types.WorkerArtifact {
	return types.WorkerArtifact{
		ID:        artifact.ID(),
		Name:      artifact.Name(),
		BuildID:   artifact.BuildID(),
		CreatedAt: artifact.CreatedAt().Unix(),
	}
}
