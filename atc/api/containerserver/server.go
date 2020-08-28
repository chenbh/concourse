package containerserver

import (
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/atc/creds"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/gc"
	"github.com/chenbh/concourse/atc/worker"
)

type Server struct {
	logger lager.Logger

	workerClient            worker.Client
	secretManager           creds.Secrets
	varSourcePool           creds.VarSourcePool
	interceptTimeoutFactory InterceptTimeoutFactory
	interceptUpdateInterval time.Duration
	containerRepository     db.ContainerRepository
	destroyer               gc.Destroyer
	clock                   clock.Clock
}

func NewServer(
	logger lager.Logger,
	workerClient worker.Client,
	secretManager creds.Secrets,
	varSourcePool creds.VarSourcePool,
	interceptTimeoutFactory InterceptTimeoutFactory,
	interceptUpdateInterval time.Duration,
	containerRepository db.ContainerRepository,
	destroyer gc.Destroyer,
	clock clock.Clock,
) *Server {
	return &Server{
		logger:                  logger,
		workerClient:            workerClient,
		secretManager:           secretManager,
		varSourcePool:           varSourcePool,
		interceptTimeoutFactory: interceptTimeoutFactory,
		interceptUpdateInterval: interceptUpdateInterval,
		containerRepository:     containerRepository,
		destroyer:               destroyer,
		clock:                   clock,
	}
}
