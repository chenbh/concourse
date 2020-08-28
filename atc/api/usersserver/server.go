package usersserver

import (
	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/atc/db"
)

type Server struct {
	logger      lager.Logger
	userFactory db.UserFactory
}

func NewServer(
	logger lager.Logger,
	userFactory db.UserFactory,
) *Server {
	return &Server{
		logger:      logger,
		userFactory: userFactory,
	}
}
