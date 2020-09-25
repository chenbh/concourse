package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) Resource(pipelineName string, resourceName string) (types.Resource, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"resource_name": resourceName,
		"team_name":     team.name,
	}

	var resource types.Resource
	err := team.connection.Send(internal.Request{
		RequestName: types.GetResource,
		Params:      params,
	}, &internal.Response{
		Result: &resource,
	})
	switch err.(type) {
	case nil:
		return resource, true, nil
	case internal.ResourceNotFoundError:
		return resource, false, nil
	default:
		return resource, false, err
	}
}

func (team *team) ListResources(pipelineName string) ([]types.Resource, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var resources []types.Resource
	err := team.connection.Send(internal.Request{
		RequestName: types.ListResources,
		Params:      params,
	}, &internal.Response{
		Result: &resources,
	})

	return resources, err
}
