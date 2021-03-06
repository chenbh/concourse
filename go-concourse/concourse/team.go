package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"io"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
)

//go:generate counterfeiter . Team

type Team interface {
	Name() string

	Auth() types.TeamAuth

	CreateOrUpdate(team types.Team) (types.Team, bool, bool, []ConfigWarning, error)
	RenameTeam(teamName, name string) (bool, []ConfigWarning, error)
	DestroyTeam(teamName string) error

	Pipeline(name string) (types.Pipeline, bool, error)
	PipelineBuilds(pipelineName string, page Page) ([]types.Build, Pagination, bool, error)
	DeletePipeline(pipelineName string) (bool, error)
	PausePipeline(pipelineName string) (bool, error)
	ArchivePipeline(pipelineName string) (bool, error)
	UnpausePipeline(pipelineName string) (bool, error)
	ExposePipeline(pipelineName string) (bool, error)
	HidePipeline(pipelineName string) (bool, error)
	RenamePipeline(pipelineName, name string) (bool, []ConfigWarning, error)
	ListPipelines() ([]types.Pipeline, error)
	PipelineConfig(pipelineName string) (types.Config, string, bool, error)
	CreateOrUpdatePipelineConfig(pipelineName string, configVersion string, passedConfig []byte, checkCredentials bool) (bool, bool, []ConfigWarning, error)

	CreatePipelineBuild(pipelineName string, plan atc.Plan) (types.Build, error)

	BuildInputsForJob(pipelineName string, jobName string) ([]types.BuildInput, bool, error)

	Job(pipelineName, jobName string) (types.Job, bool, error)
	JobBuild(pipelineName, jobName, buildName string) (types.Build, bool, error)
	JobBuilds(pipelineName string, jobName string, page Page) ([]types.Build, Pagination, bool, error)
	CreateJobBuild(pipelineName string, jobName string) (types.Build, error)
	RerunJobBuild(pipelineName string, jobName string, buildName string) (types.Build, error)
	ListJobs(pipelineName string) ([]types.Job, error)
	ScheduleJob(pipelineName string, jobName string) (bool, error)

	PauseJob(pipelineName string, jobName string) (bool, error)
	UnpauseJob(pipelineName string, jobName string) (bool, error)

	ClearTaskCache(pipelineName string, jobName string, stepName string, cachePath string) (int64, error)

	Resource(pipelineName string, resourceName string) (types.Resource, bool, error)
	ListResources(pipelineName string) ([]types.Resource, error)
	VersionedResourceTypes(pipelineName string) (types.VersionedResourceTypes, bool, error)
	ResourceVersions(pipelineName string, resourceName string, page Page, filter types.Version) ([]types.ResourceVersion, Pagination, bool, error)
	CheckResource(pipelineName string, resourceName string, version types.Version) (types.Check, bool, error)
	CheckResourceType(pipelineName string, resourceTypeName string, version types.Version) (types.Check, bool, error)
	DisableResourceVersion(pipelineName string, resourceName string, resourceVersionID int) (bool, error)
	EnableResourceVersion(pipelineName string, resourceName string, resourceVersionID int) (bool, error)

	PinResourceVersion(pipelineName string, resourceName string, resourceVersionID int) (bool, error)
	UnpinResource(pipelineName string, resourceName string) (bool, error)
	SetPinComment(pipelineName string, resourceName string, comment string) (bool, error)

	BuildsWithVersionAsInput(pipelineName string, resourceName string, resourceVersionID int) ([]types.Build, bool, error)
	BuildsWithVersionAsOutput(pipelineName string, resourceName string, resourceVersionID int) ([]types.Build, bool, error)

	ListContainers(queryList map[string]string) ([]types.Container, error)
	GetContainer(id string) (types.Container, error)
	ListVolumes() ([]types.Volume, error)
	CreateBuild(plan atc.Plan) (types.Build, error)
	Builds(page Page) ([]types.Build, Pagination, error)
	OrderingPipelines(pipelineNames []string) error

	CreateArtifact(io.Reader, string, []string) (types.WorkerArtifact, error)
	GetArtifact(int) (io.ReadCloser, error)
}

type team struct {
	name       string
	connection internal.Connection //Deprecated
	httpAgent  internal.HTTPAgent
	auth       types.TeamAuth
}

func (team *team) Name() string {
	return team.name
}

func (client *client) Team(name string) Team {
	return &team{
		name:       name,
		connection: client.connection,
	}
}

func (team *team) Auth() types.TeamAuth {
	return team.auth
}
