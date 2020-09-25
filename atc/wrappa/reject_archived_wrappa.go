package wrappa

import (
	"github.com/concourse/concourse/atc/api/pipelineserver"
	"github.com/concourse/concourse/atc/types"
	"github.com/tedsuo/rata"
)

type RejectArchivedWrappa struct {
	handlerFactory pipelineserver.RejectArchivedHandlerFactory
}

func NewRejectArchivedWrappa(factory pipelineserver.RejectArchivedHandlerFactory) *RejectArchivedWrappa {
	return &RejectArchivedWrappa{
		handlerFactory: factory,
	}
}

func (rw *RejectArchivedWrappa) Wrap(handlers rata.Handlers) rata.Handlers {
	wrapped := rata.Handlers{}

	for name, handler := range handlers {
		newHandler := handler

		switch name {
		case
			types.PausePipeline,
			types.UnpausePipeline,
			types.CreateJobBuild,
			types.ScheduleJob,
			types.CheckResource,
			types.CheckResourceType,
			types.DisableResourceVersion,
			types.EnableResourceVersion,
			types.PinResourceVersion,
			types.UnpinResource,
			types.SetPinCommentOnResource,
			types.RerunJobBuild:

			newHandler = rw.handlerFactory.RejectArchived(handler)

			// leave the handler as-is
		case
			types.GetConfig,
			types.GetBuild,
			types.BuildResources,
			types.BuildEvents,
			types.ListBuildArtifacts,
			types.GetBuildPreparation,
			types.GetBuildPlan,
			types.AbortBuild,
			types.PruneWorker,
			types.LandWorker,
			types.ReportWorkerContainers,
			types.ReportWorkerVolumes,
			types.RetireWorker,
			types.ListDestroyingContainers,
			types.ListDestroyingVolumes,
			types.GetPipeline,
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
			types.ListResources,
			types.ListResourceTypes,
			types.ListResourceVersions,
			types.GetResourceCausality,
			types.GetResourceVersion,
			types.CreateBuild,
			types.GetContainer,
			types.HijackContainer,
			types.ListContainers,
			types.ListVolumes,
			types.ListTeamBuilds,
			types.ListWorkers,
			types.RegisterWorker,
			types.HeartbeatWorker,
			types.DeleteWorker,
			types.GetTeam,
			types.SetTeam,
			types.RenameTeam,
			types.DestroyTeam,
			types.GetUser,
			types.GetInfo,
			types.GetCheck,
			types.DownloadCLI,
			types.CheckResourceWebHook,
			types.ListAllPipelines,
			types.ListBuilds,
			types.ListPipelines,
			types.ListAllJobs,
			types.ListAllResources,
			types.ListTeams,
			types.MainJobBadge,
			types.GetWall,
			types.GetLogLevel,
			types.SetLogLevel,
			types.GetInfoCreds,
			types.ListActiveUsersSince,
			types.SetWall,
			types.ClearWall,
			types.DeletePipeline,
			types.GetCC,
			types.GetVersionsDB,
			types.ListJobInputs,
			types.OrderPipelines,
			types.PauseJob,
			types.ArchivePipeline,
			types.RenamePipeline,
			types.SaveConfig,
			types.UnpauseJob,
			types.ExposePipeline,
			types.HidePipeline,
			types.CreatePipelineBuild,
			types.ClearTaskCache,
			types.CreateArtifact,
			types.GetArtifact:

		default:
			panic("how do archived pipelines affect your endpoint?")
		}

		wrapped[name] = newHandler
	}

	return wrapped
}
