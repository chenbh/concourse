package wrappa

import (
	"github.com/concourse/concourse/atc/api/auth"
	"github.com/concourse/concourse/atc/types"
	"github.com/tedsuo/rata"
)

type APIAuthWrappa struct {
	checkPipelineAccessHandlerFactory   auth.CheckPipelineAccessHandlerFactory
	checkBuildReadAccessHandlerFactory  auth.CheckBuildReadAccessHandlerFactory
	checkBuildWriteAccessHandlerFactory auth.CheckBuildWriteAccessHandlerFactory
	checkWorkerTeamAccessHandlerFactory auth.CheckWorkerTeamAccessHandlerFactory
}

func NewAPIAuthWrappa(
	checkPipelineAccessHandlerFactory auth.CheckPipelineAccessHandlerFactory,
	checkBuildReadAccessHandlerFactory auth.CheckBuildReadAccessHandlerFactory,
	checkBuildWriteAccessHandlerFactory auth.CheckBuildWriteAccessHandlerFactory,
	checkWorkerTeamAccessHandlerFactory auth.CheckWorkerTeamAccessHandlerFactory,
) *APIAuthWrappa {
	return &APIAuthWrappa{
		checkPipelineAccessHandlerFactory:   checkPipelineAccessHandlerFactory,
		checkBuildReadAccessHandlerFactory:  checkBuildReadAccessHandlerFactory,
		checkBuildWriteAccessHandlerFactory: checkBuildWriteAccessHandlerFactory,
		checkWorkerTeamAccessHandlerFactory: checkWorkerTeamAccessHandlerFactory,
	}
}

func (wrappa *APIAuthWrappa) Wrap(handlers rata.Handlers) rata.Handlers {
	wrapped := rata.Handlers{}

	rejector := auth.UnauthorizedRejector{}

	for name, handler := range handlers {
		newHandler := handler

		switch name {
		// pipeline is public or authorized
		case types.GetBuild,
			types.BuildResources:
			newHandler = wrappa.checkBuildReadAccessHandlerFactory.AnyJobHandler(handler, rejector)

		// pipeline and job are public or authorized
		case types.GetBuildPreparation,
			types.BuildEvents,
			types.GetBuildPlan,
			types.ListBuildArtifacts:
			newHandler = wrappa.checkBuildReadAccessHandlerFactory.CheckIfPrivateJobHandler(handler, rejector)

			// resource belongs to authorized team
		case types.AbortBuild:
			newHandler = wrappa.checkBuildWriteAccessHandlerFactory.HandlerFor(handler, rejector)

		// requester is system, admin team, or worker owning team
		case types.PruneWorker,
			types.LandWorker,
			types.RetireWorker,
			types.ListDestroyingVolumes,
			types.ListDestroyingContainers,
			types.ReportWorkerContainers,
			types.ReportWorkerVolumes:
			newHandler = wrappa.checkWorkerTeamAccessHandlerFactory.HandlerFor(handler, rejector)

		// pipeline is public or authorized
		case types.GetPipeline,
			types.GetJobBuild,
			types.PipelineBadge,
			types.JobBadge,
			types.ListJobs,
			types.GetJob,
			types.ListJobBuilds,
			types.ListPipelineBuilds,
			types.GetResource,
			types.ListBuildsWithVersionAsInput,
			types.ListBuildsWithVersionAsOutput,
			types.GetResourceCausality,
			types.GetResourceVersion,
			types.ListResources,
			types.ListResourceTypes,
			types.ListResourceVersions:
			newHandler = wrappa.checkPipelineAccessHandlerFactory.HandlerFor(handler, rejector)

		// authenticated
		case types.CreateBuild,
			types.GetContainer,
			types.HijackContainer,
			types.ListContainers,
			types.ListWorkers,
			types.RegisterWorker,
			types.HeartbeatWorker,
			types.DeleteWorker,
			types.GetTeam,
			types.SetTeam,
			types.ListTeamBuilds,
			types.RenameTeam,
			types.DestroyTeam,
			types.ListVolumes,
			types.GetUser:
			newHandler = auth.CheckAuthenticationHandler(handler, rejector)

		// unauthenticated / delegating to handler (validate token if provided)
		case types.DownloadCLI,
			types.CheckResourceWebHook,
			types.GetInfo,
			types.GetCheck,
			types.ListTeams,
			types.ListAllPipelines,
			types.ListPipelines,
			types.ListAllJobs,
			types.ListAllResources,
			types.ListBuilds,
			types.MainJobBadge,
			types.GetWall:
			newHandler = auth.CheckAuthenticationIfProvidedHandler(handler, rejector)

		case types.GetLogLevel,
			types.ListActiveUsersSince,
			types.SetLogLevel,
			types.GetInfoCreds,
			types.SetWall,
			types.ClearWall:
			newHandler = auth.CheckAdminHandler(handler, rejector)

		// authorized (requested team matches resource team)
		case types.CheckResource,
			types.CheckResourceType,
			types.CreateJobBuild,
			types.RerunJobBuild,
			types.CreatePipelineBuild,
			types.DeletePipeline,
			types.DisableResourceVersion,
			types.EnableResourceVersion,
			types.PinResourceVersion,
			types.UnpinResource,
			types.SetPinCommentOnResource,
			types.GetConfig,
			types.GetCC,
			types.GetVersionsDB,
			types.ListJobInputs,
			types.OrderPipelines,
			types.PauseJob,
			types.PausePipeline,
			types.RenamePipeline,
			types.UnpauseJob,
			types.UnpausePipeline,
			types.ExposePipeline,
			types.HidePipeline,
			types.SaveConfig,
			types.ArchivePipeline,
			types.ClearTaskCache,
			types.CreateArtifact,
			types.ScheduleJob,
			types.GetArtifact:
			newHandler = auth.CheckAuthorizationHandler(handler, rejector)

		// think about it!
		default:
			panic("you missed a spot")
		}

		wrapped[name] = newHandler
	}

	return wrapped
}
