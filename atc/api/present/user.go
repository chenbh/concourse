package present

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
)

func User(user db.User) atc.User {
	return atc.User{
		ID:        user.ID(),
		Username:  user.Name(),
		Connector: user.Connector(),
		LastLogin: user.LastLogin().Unix(),
	}
}
