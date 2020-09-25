package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) ListVolumes() ([]types.Volume, error) {
	var volumes []types.Volume

	params := rata.Params{
		"team_name": team.name,
	}
	err := team.connection.Send(internal.Request{
		RequestName: types.ListVolumes,
		Params:      params,
	}, &internal.Response{
		Result: &volumes,
	})

	return volumes, err
}
