package auditor

import (
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter . Auditor

func NewAuditor(
	EnableBuildAuditLog bool,
	EnableContainerAuditLog bool,
	EnableJobAuditLog bool,
	EnablePipelineAuditLog bool,
	EnableResourceAuditLog bool,
	EnableSystemAuditLog bool,
	EnableTeamAuditLog bool,
	EnableWorkerAuditLog bool,
	EnableVolumeAuditLog bool,
	logger lager.Logger,
) *auditor {
	return &auditor{
		EnableBuildAuditLog:     EnableBuildAuditLog,
		EnableContainerAuditLog: EnableContainerAuditLog,
		EnableJobAuditLog:       EnableJobAuditLog,
		EnablePipelineAuditLog:  EnablePipelineAuditLog,
		EnableResourceAuditLog:  EnableResourceAuditLog,
		EnableSystemAuditLog:    EnableSystemAuditLog,
		EnableTeamAuditLog:      EnableTeamAuditLog,
		EnableWorkerAuditLog:    EnableWorkerAuditLog,
		EnableVolumeAuditLog:    EnableVolumeAuditLog,
		logger:                  logger,
	}
}

type Auditor interface {
	Audit(action string, userName string, r *http.Request)
}

type auditor struct {
	EnableBuildAuditLog     bool
	EnableContainerAuditLog bool
	EnableJobAuditLog       bool
	EnablePipelineAuditLog  bool
	EnableResourceAuditLog  bool
	EnableSystemAuditLog    bool
	EnableTeamAuditLog      bool
	EnableWorkerAuditLog    bool
	EnableVolumeAuditLog    bool
	logger                  lager.Logger
}

func (a *auditor) ValidateAction(action string) bool {
	switch action {
	case types.GetBuild,
		types.GetBuildPlan,
		types.CreateBuild,
		types.RerunJobBuild,
		types.ListBuilds,
		types.BuildEvents,
		types.BuildResources,
		types.AbortBuild,
		types.GetBuildPreparation,
		types.ListBuildsWithVersionAsInput,
		types.ListBuildsWithVersionAsOutput,
		types.CreateArtifact,
		types.GetArtifact,
		types.ListBuildArtifacts:
		return a.EnableBuildAuditLog
	case types.ListContainers,
		types.GetContainer,
		types.HijackContainer,
		types.ListDestroyingContainers,
		types.ReportWorkerContainers:
		return a.EnableContainerAuditLog
	case types.GetJob,
		types.CreateJobBuild,
		types.ListAllJobs,
		types.ListJobs,
		types.ListJobBuilds,
		types.ListJobInputs,
		types.GetJobBuild,
		types.PauseJob,
		types.UnpauseJob,
		types.ScheduleJob,
		types.JobBadge,
		types.MainJobBadge:
		return a.EnableJobAuditLog
	case types.ListAllPipelines,
		types.ListPipelines,
		types.GetPipeline,
		types.DeletePipeline,
		types.OrderPipelines,
		types.PausePipeline,
		types.ArchivePipeline,
		types.UnpausePipeline,
		types.ExposePipeline,
		types.HidePipeline,
		types.RenamePipeline,
		types.ListPipelineBuilds,
		types.CreatePipelineBuild,
		types.PipelineBadge:
		return a.EnablePipelineAuditLog
	case types.ListAllResources,
		types.ListResources,
		types.ListResourceTypes,
		types.GetResource,
		types.UnpinResource,
		types.SetPinCommentOnResource,
		types.CheckResource,
		types.CheckResourceWebHook,
		types.CheckResourceType,
		types.ListResourceVersions,
		types.GetResourceVersion,
		types.EnableResourceVersion,
		types.DisableResourceVersion,
		types.PinResourceVersion,
		types.GetResourceCausality,
		types.GetCheck:
		return a.EnableResourceAuditLog
	case
		types.SaveConfig,
		types.GetConfig,
		types.GetCC,
		types.GetVersionsDB,
		types.ClearTaskCache,
		types.SetLogLevel,
		types.GetLogLevel,
		types.DownloadCLI,
		types.GetInfo,
		types.GetInfoCreds,
		types.ListActiveUsersSince,
		types.GetUser,
		types.GetWall,
		types.SetWall,
		types.ClearWall:
		return a.EnableSystemAuditLog
	case types.ListTeams,
		types.SetTeam,
		types.RenameTeam,
		types.DestroyTeam,
		types.ListTeamBuilds,
		types.GetTeam:
		return a.EnableTeamAuditLog
	case types.RegisterWorker,
		types.LandWorker,
		types.RetireWorker,
		types.PruneWorker,
		types.HeartbeatWorker,
		types.ListWorkers,
		types.DeleteWorker:
		return a.EnableWorkerAuditLog
	case types.ListVolumes,
		types.ListDestroyingVolumes,
		types.ReportWorkerVolumes:
		return a.EnableVolumeAuditLog
	default:
		panic(fmt.Sprintf("unhandled action: %s", action))
	}
}

func (a *auditor) Audit(action string, userName string, r *http.Request) {
	err := r.ParseForm()
	if err == nil && a.ValidateAction(action) {
		a.logger.Info("audit", lager.Data{"action": action, "user": userName, "parameters": r.Form})
	}
}
