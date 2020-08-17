package present

import (
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/db"
)

func Team(team db.Team) atc.Team {
	return atc.Team{
		ID:   team.ID(),
		Name: team.Name(),
		Auth: team.Auth(),
	}
}
