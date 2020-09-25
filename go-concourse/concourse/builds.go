package concourse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) CreateBuild(plan atc.Plan) (types.Build, error) {
	var build types.Build

	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(plan)
	if err != nil {
		return build, fmt.Errorf("Unable to marshal plan: %s", err)
	}
	err = team.connection.Send(internal.Request{
		RequestName: types.CreateBuild,
		Body:        buffer,
		Params: rata.Params{
			"team_name": team.Name(),
		},
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
	}, &internal.Response{
		Result: &build,
	})

	return build, err
}

func (team *team) CreateJobBuild(pipelineName string, jobName string) (types.Build, error) {
	params := rata.Params{
		"job_name":      jobName,
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var build types.Build
	err := team.connection.Send(internal.Request{
		RequestName: types.CreateJobBuild,
		Params:      params,
	}, &internal.Response{
		Result: &build,
	})

	return build, err
}

func (team *team) RerunJobBuild(pipelineName string, jobName string, buildName string) (types.Build, error) {
	params := rata.Params{
		"build_name":    buildName,
		"job_name":      jobName,
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var build types.Build
	err := team.connection.Send(internal.Request{
		RequestName: types.RerunJobBuild,
		Params:      params,
	}, &internal.Response{
		Result: &build,
	})

	return build, err
}

func (team *team) JobBuild(pipelineName, jobName, buildName string) (types.Build, bool, error) {
	params := rata.Params{
		"job_name":      jobName,
		"build_name":    buildName,
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var build types.Build
	err := team.connection.Send(internal.Request{
		RequestName: types.GetJobBuild,
		Params:      params,
	}, &internal.Response{
		Result: &build,
	})

	switch err.(type) {
	case nil:
		return build, true, nil
	case internal.ResourceNotFoundError:
		return build, false, nil
	default:
		return build, false, err
	}
}

func (client *client) Build(buildID string) (types.Build, bool, error) {
	params := rata.Params{
		"build_id": buildID,
	}

	var build types.Build
	err := client.connection.Send(internal.Request{
		RequestName: types.GetBuild,
		Params:      params,
	}, &internal.Response{
		Result: &build,
	})

	switch err.(type) {
	case nil:
		return build, true, nil
	case internal.ResourceNotFoundError:
		return build, false, nil
	default:
		return build, false, err
	}
}

func (client *client) Builds(page Page) ([]types.Build, Pagination, error) {
	var builds []types.Build

	headers := http.Header{}
	err := client.connection.Send(internal.Request{
		RequestName: types.ListBuilds,
		Query:       page.QueryParams(),
	}, &internal.Response{
		Result:  &builds,
		Headers: &headers,
	})

	switch err.(type) {
	case nil:
		pagination, err := paginationFromHeaders(headers)
		if err != nil {
			return nil, Pagination{}, err
		}

		return builds, pagination, nil
	default:
		return nil, Pagination{}, err
	}
}

func (client *client) AbortBuild(buildID string) error {
	params := rata.Params{
		"build_id": buildID,
	}

	return client.connection.Send(internal.Request{
		RequestName: types.AbortBuild,
		Params:      params,
	}, nil)
}

func (team *team) Builds(page Page) ([]types.Build, Pagination, error) {
	var builds []types.Build

	headers := http.Header{}

	params := rata.Params{
		"team_name": team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: types.ListTeamBuilds,
		Params:      params,
		Query:       page.QueryParams(),
	}, &internal.Response{
		Result:  &builds,
		Headers: &headers,
	})

	switch err.(type) {
	case nil:
		pagination, err := paginationFromHeaders(headers)
		if err != nil {
			return nil, Pagination{}, err
		}

		return builds, pagination, nil
	default:
		return nil, Pagination{}, err
	}
}

func (client *client) ListBuildArtifacts(buildID string) ([]types.WorkerArtifact, error) {
	params := rata.Params{
		"build_id": buildID,
	}

	var artifacts []types.WorkerArtifact

	err := client.connection.Send(internal.Request{
		RequestName: types.ListBuildArtifacts,
		Params:      params,
	}, &internal.Response{
		Result: &artifacts,
	})

	return artifacts, err
}
