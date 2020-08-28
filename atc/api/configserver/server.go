package configserver

import (
	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/atc/creds"
	"github.com/chenbh/concourse/atc/db"
)

type Server struct {
	logger        lager.Logger
	teamFactory   db.TeamFactory
	secretManager creds.Secrets
}

func NewServer(
	logger lager.Logger,
	teamFactory db.TeamFactory,
	secretManager creds.Secrets,
) *Server {
	return &Server{
		logger:        logger,
		teamFactory:   teamFactory,
		secretManager: secretManager,
	}
}
