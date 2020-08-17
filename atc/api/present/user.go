package present

import (
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/db"
)

func User(user db.User) atc.User {
	return atc.User{
		ID:        user.ID(),
		Username:  user.Name(),
		Connector: user.Connector(),
		LastLogin: user.LastLogin().Unix(),
	}
}
