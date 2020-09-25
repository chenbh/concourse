package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"strconv"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (client *client) BuildResources(buildID int) (types.BuildInputsOutputs, bool, error) {
	params := rata.Params{
		"build_id": strconv.Itoa(buildID),
	}

	var buildInputsOutputs types.BuildInputsOutputs
	err := client.connection.Send(internal.Request{
		RequestName: types.BuildResources,
		Params:      params,
	}, &internal.Response{
		Result: &buildInputsOutputs,
	})

	switch err.(type) {
	case nil:
		return buildInputsOutputs, true, nil
	case internal.ResourceNotFoundError:
		return buildInputsOutputs, false, nil
	default:
		return buildInputsOutputs, false, err
	}
}
