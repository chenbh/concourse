package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) VersionedResourceTypes(pipelineName string) (types.VersionedResourceTypes, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var versionedResourceTypes types.VersionedResourceTypes
	err := team.connection.Send(internal.Request{
		RequestName: types.ListResourceTypes,
		Params:      params,
	}, &internal.Response{
		Result: &versionedResourceTypes,
	})

	switch err.(type) {
	case nil:
		return versionedResourceTypes, true, nil
	case internal.ResourceNotFoundError:
		return versionedResourceTypes, false, nil
	default:
		return versionedResourceTypes, false, err
	}
}
