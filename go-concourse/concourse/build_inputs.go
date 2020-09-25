package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"strconv"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) BuildInputsForJob(pipelineName string, jobName string) ([]types.BuildInput, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	var buildInputs []types.BuildInput
	err := team.connection.Send(internal.Request{
		RequestName: types.ListJobInputs,
		Params:      params,
	}, &internal.Response{
		Result: &buildInputs,
	})

	switch err.(type) {
	case nil:
		return buildInputs, true, nil
	case internal.ResourceNotFoundError:
		return buildInputs, false, nil
	default:
		return buildInputs, false, err
	}
}

func (team *team) BuildsWithVersionAsInput(pipelineName string, resourceName string, resourceVersionID int) ([]types.Build, bool, error) {
	params := rata.Params{
		"pipeline_name":              pipelineName,
		"resource_name":              resourceName,
		"resource_config_version_id": strconv.Itoa(resourceVersionID),
		"team_name":                  team.name,
	}

	var builds []types.Build
	err := team.connection.Send(internal.Request{
		RequestName: types.ListBuildsWithVersionAsInput,
		Params:      params,
	}, &internal.Response{
		Result: &builds,
	})

	switch err.(type) {
	case nil:
		return builds, true, nil
	case internal.ResourceNotFoundError:
		return builds, false, nil
	default:
		return builds, false, err
	}
}
