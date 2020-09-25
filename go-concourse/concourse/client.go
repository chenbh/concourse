package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"io"
	"net/http"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
)

//go:generate counterfeiter . Client

type Client interface {
	URL() string
	HTTPClient() *http.Client
	Builds(Page) ([]types.Build, Pagination, error)
	Build(buildID string) (types.Build, bool, error)
	BuildEvents(buildID string) (Events, error)
	BuildResources(buildID int) (types.BuildInputsOutputs, bool, error)
	ListBuildArtifacts(buildID string) ([]types.WorkerArtifact, error)
	AbortBuild(buildID string) error
	BuildPlan(buildID int) (atc.PublicBuildPlan, bool, error)
	SaveWorker(types.Worker, *time.Duration) (*types.Worker, error)
	ListWorkers() ([]types.Worker, error)
	PruneWorker(workerName string) error
	LandWorker(workerName string) error
	GetInfo() (types.Info, error)
	GetCLIReader(arch, platform string) (io.ReadCloser, http.Header, error)
	ListPipelines() ([]types.Pipeline, error)
	ListAllJobs() ([]types.Job, error)
	ListTeams() ([]types.Team, error)
	FindTeam(teamName string) (Team, error)
	Team(teamName string) Team
	UserInfo() (types.UserInfo, error)
	ListActiveUsersSince(since time.Time) ([]types.User, error)
	Check(checkID string) (types.Check, bool, error)
}

type client struct {
	connection internal.Connection //Deprecated
	httpAgent  internal.HTTPAgent
}

func NewClient(apiURL string, httpClient *http.Client, tracing bool) Client {
	return &client{
		connection: internal.NewConnection(apiURL, httpClient, tracing),
		httpAgent:  internal.NewHTTPAgent(apiURL, httpClient, tracing),
	}
}

func (client *client) URL() string {
	return client.connection.URL()
}

func (client *client) HTTPClient() *http.Client {
	return client.connection.HTTPClient()
}
