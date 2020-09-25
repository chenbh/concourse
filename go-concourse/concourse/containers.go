package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"net/url"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) ListContainers(queryList map[string]string) ([]types.Container, error) {
	var containers []types.Container
	urlValues := url.Values{}

	params := rata.Params{
		"team_name": team.name,
	}
	for k, v := range queryList {
		urlValues[k] = []string{v}
	}
	err := team.connection.Send(internal.Request{
		RequestName: types.ListContainers,
		Query:       urlValues,
		Params:      params,
	}, &internal.Response{
		Result: &containers,
	})
	return containers, err
}

func (team *team) GetContainer(handle string) (types.Container, error) {
	var container types.Container

	params := rata.Params{
		"id":        handle,
		"team_name": team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: types.GetContainer,
		Params:      params,
	}, &internal.Response{
		Result: &container,
	})

	return container, err
}
