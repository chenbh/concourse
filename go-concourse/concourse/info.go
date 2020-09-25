package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
)

func (client *client) GetInfo() (types.Info, error) {
	var info types.Info

	err := client.connection.Send(internal.Request{
		RequestName: types.GetInfo,
	}, &internal.Response{
		Result: &info,
	})

	return info, err
}
