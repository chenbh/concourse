package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Team(team db.Team) types.Team {
	return types.Team{
		ID:   team.ID(),
		Name: team.Name(),
		Auth: team.Auth(),
	}
}
