package concourse

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/go-concourse/concourse/internal"
)

func (client *client) GetInfo() (atc.Info, error) {
	var info atc.Info

	err := client.connection.Send(internal.Request{
		RequestName: atc.GetInfo,
	}, &internal.Response{
		Result: &info,
	})

	return info, err
}
