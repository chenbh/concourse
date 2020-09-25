package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/go-concourse/concourse/eventstream"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

type Events interface {
	NextEvent() (types.Event, error)
	Close() error
}

func (client *client) BuildEvents(buildID string) (Events, error) {
	sseEvents, err := client.connection.ConnectToEventStream(internal.Request{
		RequestName: types.BuildEvents,
		Params: rata.Params{
			"build_id": buildID,
		},
	})
	if err != nil {
		return nil, err
	}

	return eventstream.NewSSEEventStream(sseEvents), nil
}
