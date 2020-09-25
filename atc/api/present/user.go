package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func User(user db.User) types.User {
	return types.User{
		ID:        user.ID(),
		Username:  user.Name(),
		Connector: user.Connector(),
		LastLogin: user.LastLogin().Unix(),
	}
}
