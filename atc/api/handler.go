package api

import (
	"github.com/concourse/concourse/atc/types"
	"net/http"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/api/artifactserver"
	"github.com/concourse/concourse/atc/api/buildserver"
	"github.com/concourse/concourse/atc/api/ccserver"
	"github.com/concourse/concourse/atc/api/checkserver"
	"github.com/concourse/concourse/atc/api/cliserver"
	"github.com/concourse/concourse/atc/api/configserver"
	"github.com/concourse/concourse/atc/api/containerserver"
	"github.com/concourse/concourse/atc/api/infoserver"
	"github.com/concourse/concourse/atc/api/jobserver"
	"github.com/concourse/concourse/atc/api/loglevelserver"
	"github.com/concourse/concourse/atc/api/pipelineserver"
	"github.com/concourse/concourse/atc/api/resourceserver"
	"github.com/concourse/concourse/atc/api/resourceserver/versionserver"
	"github.com/concourse/concourse/atc/api/teamserver"
	"github.com/concourse/concourse/atc/api/usersserver"
	"github.com/concourse/concourse/atc/api/volumeserver"
	"github.com/concourse/concourse/atc/api/wallserver"
	"github.com/concourse/concourse/atc/api/workerserver"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/gc"
	"github.com/concourse/concourse/atc/mainredirect"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/atc/wrappa"
	"github.com/tedsuo/rata"
)

func NewHandler(
	logger lager.Logger,

	externalURL string,
	clusterName string,

	wrapper wrappa.Wrappa,

	dbTeamFactory db.TeamFactory,
	dbPipelineFactory db.PipelineFactory,
	dbJobFactory db.JobFactory,
	dbResourceFactory db.ResourceFactory,
	dbWorkerFactory db.WorkerFactory,
	volumeRepository db.VolumeRepository,
	containerRepository db.ContainerRepository,
	destroyer gc.Destroyer,
	dbBuildFactory db.BuildFactory,
	dbCheckFactory db.CheckFactory,
	dbResourceConfigFactory db.ResourceConfigFactory,
	dbUserFactory db.UserFactory,

	eventHandlerFactory buildserver.EventHandlerFactory,

	workerClient worker.Client,

	sink *lager.ReconfigurableSink,

	isTLSEnabled bool,

	cliDownloadsDir string,
	version string,
	workerVersion string,
	secretManager creds.Secrets,
	varSourcePool creds.VarSourcePool,
	credsManagers creds.Managers,
	interceptTimeoutFactory containerserver.InterceptTimeoutFactory,
	interceptUpdateInterval time.Duration,
	dbWall db.Wall,
	clock clock.Clock,
) (http.Handler, error) {

	absCLIDownloadsDir, err := filepath.Abs(cliDownloadsDir)
	if err != nil {
		return nil, err
	}

	pipelineHandlerFactory := pipelineserver.NewScopedHandlerFactory(dbTeamFactory)
	buildHandlerFactory := buildserver.NewScopedHandlerFactory(logger)
	teamHandlerFactory := NewTeamScopedHandlerFactory(logger, dbTeamFactory)

	buildServer := buildserver.NewServer(logger, externalURL, dbTeamFactory, dbBuildFactory, eventHandlerFactory)
	checkServer := checkserver.NewServer(logger, dbCheckFactory)
	jobServer := jobserver.NewServer(logger, externalURL, secretManager, dbJobFactory, dbCheckFactory)
	resourceServer := resourceserver.NewServer(logger, secretManager, varSourcePool, dbCheckFactory, dbResourceFactory, dbResourceConfigFactory)

	versionServer := versionserver.NewServer(logger, externalURL)
	pipelineServer := pipelineserver.NewServer(logger, dbTeamFactory, dbPipelineFactory, externalURL)
	configServer := configserver.NewServer(logger, dbTeamFactory, secretManager)
	ccServer := ccserver.NewServer(logger, dbTeamFactory, externalURL)
	workerServer := workerserver.NewServer(logger, dbTeamFactory, dbWorkerFactory)
	logLevelServer := loglevelserver.NewServer(logger, sink)
	cliServer := cliserver.NewServer(logger, absCLIDownloadsDir)
	containerServer := containerserver.NewServer(logger, workerClient, secretManager, varSourcePool, interceptTimeoutFactory, interceptUpdateInterval, containerRepository, destroyer, clock)
	volumesServer := volumeserver.NewServer(logger, volumeRepository, destroyer)
	teamServer := teamserver.NewServer(logger, dbTeamFactory, externalURL)
	infoServer := infoserver.NewServer(logger, version, workerVersion, externalURL, clusterName, credsManagers)
	artifactServer := artifactserver.NewServer(logger, workerClient)
	usersServer := usersserver.NewServer(logger, dbUserFactory)
	wallServer := wallserver.NewServer(dbWall, logger)

	handlers := map[string]http.Handler{
		types.GetConfig:  http.HandlerFunc(configServer.GetConfig),
		types.SaveConfig: http.HandlerFunc(configServer.SaveConfig),

		types.GetCC: http.HandlerFunc(ccServer.GetCC),

		types.ListBuilds:          http.HandlerFunc(buildServer.ListBuilds),
		types.CreateBuild:         teamHandlerFactory.HandlerFor(buildServer.CreateBuild),
		types.GetBuild:            buildHandlerFactory.HandlerFor(buildServer.GetBuild),
		types.BuildResources:      buildHandlerFactory.HandlerFor(buildServer.BuildResources),
		types.AbortBuild:          buildHandlerFactory.HandlerFor(buildServer.AbortBuild),
		types.GetBuildPlan:        buildHandlerFactory.HandlerFor(buildServer.GetBuildPlan),
		types.GetBuildPreparation: buildHandlerFactory.HandlerFor(buildServer.GetBuildPreparation),
		types.BuildEvents:         buildHandlerFactory.HandlerFor(buildServer.BuildEvents),
		types.ListBuildArtifacts:  buildHandlerFactory.HandlerFor(buildServer.GetBuildArtifacts),

		types.GetCheck: http.HandlerFunc(checkServer.GetCheck),

		types.ListAllJobs:    http.HandlerFunc(jobServer.ListAllJobs),
		types.ListJobs:       pipelineHandlerFactory.HandlerFor(jobServer.ListJobs),
		types.GetJob:         pipelineHandlerFactory.HandlerFor(jobServer.GetJob),
		types.ListJobBuilds:  pipelineHandlerFactory.HandlerFor(jobServer.ListJobBuilds),
		types.ListJobInputs:  pipelineHandlerFactory.HandlerFor(jobServer.ListJobInputs),
		types.GetJobBuild:    pipelineHandlerFactory.HandlerFor(jobServer.GetJobBuild),
		types.CreateJobBuild: pipelineHandlerFactory.HandlerFor(jobServer.CreateJobBuild),
		types.RerunJobBuild:  pipelineHandlerFactory.HandlerFor(jobServer.RerunJobBuild),
		types.PauseJob:       pipelineHandlerFactory.HandlerFor(jobServer.PauseJob),
		types.UnpauseJob:     pipelineHandlerFactory.HandlerFor(jobServer.UnpauseJob),
		types.ScheduleJob:    pipelineHandlerFactory.HandlerFor(jobServer.ScheduleJob),
		types.JobBadge:       pipelineHandlerFactory.HandlerFor(jobServer.JobBadge),
		types.MainJobBadge: mainredirect.Handler{
			Routes: types.Routes,
			Route:  types.JobBadge,
		},

		types.ClearTaskCache: pipelineHandlerFactory.HandlerFor(jobServer.ClearTaskCache),

		types.ListAllPipelines:    http.HandlerFunc(pipelineServer.ListAllPipelines),
		types.ListPipelines:       http.HandlerFunc(pipelineServer.ListPipelines),
		types.GetPipeline:         pipelineHandlerFactory.HandlerFor(pipelineServer.GetPipeline),
		types.DeletePipeline:      pipelineHandlerFactory.HandlerFor(pipelineServer.DeletePipeline),
		types.OrderPipelines:      http.HandlerFunc(pipelineServer.OrderPipelines),
		types.PausePipeline:       pipelineHandlerFactory.HandlerFor(pipelineServer.PausePipeline),
		types.ArchivePipeline:     pipelineHandlerFactory.HandlerFor(pipelineServer.ArchivePipeline),
		types.UnpausePipeline:     pipelineHandlerFactory.HandlerFor(pipelineServer.UnpausePipeline),
		types.ExposePipeline:      pipelineHandlerFactory.HandlerFor(pipelineServer.ExposePipeline),
		types.HidePipeline:        pipelineHandlerFactory.HandlerFor(pipelineServer.HidePipeline),
		types.GetVersionsDB:       pipelineHandlerFactory.HandlerFor(pipelineServer.GetVersionsDB),
		types.RenamePipeline:      pipelineHandlerFactory.HandlerFor(pipelineServer.RenamePipeline),
		types.ListPipelineBuilds:  pipelineHandlerFactory.HandlerFor(pipelineServer.ListPipelineBuilds),
		types.CreatePipelineBuild: pipelineHandlerFactory.HandlerFor(pipelineServer.CreateBuild),
		types.PipelineBadge:       pipelineHandlerFactory.HandlerFor(pipelineServer.PipelineBadge),

		types.ListAllResources:        http.HandlerFunc(resourceServer.ListAllResources),
		types.ListResources:           pipelineHandlerFactory.HandlerFor(resourceServer.ListResources),
		types.ListResourceTypes:       pipelineHandlerFactory.HandlerFor(resourceServer.ListVersionedResourceTypes),
		types.GetResource:             pipelineHandlerFactory.HandlerFor(resourceServer.GetResource),
		types.UnpinResource:           pipelineHandlerFactory.HandlerFor(resourceServer.UnpinResource),
		types.SetPinCommentOnResource: pipelineHandlerFactory.HandlerFor(resourceServer.SetPinCommentOnResource),
		types.CheckResource:           pipelineHandlerFactory.HandlerFor(resourceServer.CheckResource),
		types.CheckResourceWebHook:    pipelineHandlerFactory.HandlerFor(resourceServer.CheckResourceWebHook),
		types.CheckResourceType:       pipelineHandlerFactory.HandlerFor(resourceServer.CheckResourceType),

		types.ListResourceVersions:          pipelineHandlerFactory.HandlerFor(versionServer.ListResourceVersions),
		types.GetResourceVersion:            pipelineHandlerFactory.HandlerFor(versionServer.GetResourceVersion),
		types.EnableResourceVersion:         pipelineHandlerFactory.HandlerFor(versionServer.EnableResourceVersion),
		types.DisableResourceVersion:        pipelineHandlerFactory.HandlerFor(versionServer.DisableResourceVersion),
		types.PinResourceVersion:            pipelineHandlerFactory.HandlerFor(versionServer.PinResourceVersion),
		types.ListBuildsWithVersionAsInput:  pipelineHandlerFactory.HandlerFor(versionServer.ListBuildsWithVersionAsInput),
		types.ListBuildsWithVersionAsOutput: pipelineHandlerFactory.HandlerFor(versionServer.ListBuildsWithVersionAsOutput),
		types.GetResourceCausality:          pipelineHandlerFactory.HandlerFor(versionServer.GetCausality),

		types.ListWorkers:     http.HandlerFunc(workerServer.ListWorkers),
		types.RegisterWorker:  http.HandlerFunc(workerServer.RegisterWorker),
		types.LandWorker:      http.HandlerFunc(workerServer.LandWorker),
		types.RetireWorker:    http.HandlerFunc(workerServer.RetireWorker),
		types.PruneWorker:     http.HandlerFunc(workerServer.PruneWorker),
		types.HeartbeatWorker: http.HandlerFunc(workerServer.HeartbeatWorker),
		types.DeleteWorker:    http.HandlerFunc(workerServer.DeleteWorker),

		types.SetLogLevel: http.HandlerFunc(logLevelServer.SetMinLevel),
		types.GetLogLevel: http.HandlerFunc(logLevelServer.GetMinLevel),

		types.DownloadCLI:  http.HandlerFunc(cliServer.Download),
		types.GetInfo:      http.HandlerFunc(infoServer.Info),
		types.GetInfoCreds: http.HandlerFunc(infoServer.Creds),

		types.GetUser:              http.HandlerFunc(usersServer.GetUser),
		types.ListActiveUsersSince: http.HandlerFunc(usersServer.GetUsersSince),

		types.ListContainers:           teamHandlerFactory.HandlerFor(containerServer.ListContainers),
		types.GetContainer:             teamHandlerFactory.HandlerFor(containerServer.GetContainer),
		types.HijackContainer:          teamHandlerFactory.HandlerFor(containerServer.HijackContainer),
		types.ListDestroyingContainers: http.HandlerFunc(containerServer.ListDestroyingContainers),
		types.ReportWorkerContainers:   http.HandlerFunc(containerServer.ReportWorkerContainers),

		types.ListVolumes:           teamHandlerFactory.HandlerFor(volumesServer.ListVolumes),
		types.ListDestroyingVolumes: http.HandlerFunc(volumesServer.ListDestroyingVolumes),
		types.ReportWorkerVolumes:   http.HandlerFunc(volumesServer.ReportWorkerVolumes),

		types.ListTeams:      http.HandlerFunc(teamServer.ListTeams),
		types.GetTeam:        http.HandlerFunc(teamServer.GetTeam),
		types.SetTeam:        http.HandlerFunc(teamServer.SetTeam),
		types.RenameTeam:     http.HandlerFunc(teamServer.RenameTeam),
		types.DestroyTeam:    http.HandlerFunc(teamServer.DestroyTeam),
		types.ListTeamBuilds: http.HandlerFunc(teamServer.ListTeamBuilds),

		types.CreateArtifact: teamHandlerFactory.HandlerFor(artifactServer.CreateArtifact),
		types.GetArtifact:    teamHandlerFactory.HandlerFor(artifactServer.GetArtifact),

		types.GetWall:   http.HandlerFunc(wallServer.GetWall),
		types.SetWall:   http.HandlerFunc(wallServer.SetWall),
		types.ClearWall: http.HandlerFunc(wallServer.ClearWall),
	}

	return rata.NewRouter(types.Routes, wrapper.Wrap(handlers))
}
