package concourse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"net/http"
	"time"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

type PruneWorkerError struct {
	types.PruneWorkerResponseBody
}

func (e PruneWorkerError) Error() string {
	return e.Stderr
}

func (client *client) ListWorkers() ([]types.Worker, error) {
	var workers []types.Worker
	err := client.connection.Send(internal.Request{
		RequestName: types.ListWorkers,
	}, &internal.Response{
		Result: &workers,
	})
	return workers, err
}

func (client *client) SaveWorker(worker types.Worker, ttl *time.Duration) (*types.Worker, error) {
	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(worker)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal worker: %s", err)
	}

	params := rata.Params{}
	if ttl != nil {
		params["ttl"] = ttl.String()
	}

	var savedWorker *types.Worker
	err = client.connection.Send(internal.Request{
		RequestName: types.RegisterWorker,
		Body:        buffer,
		Params:      params,
	}, &internal.Response{
		Result: &savedWorker,
	})

	return savedWorker, err
}

func (client *client) PruneWorker(workerName string) error {
	params := rata.Params{"worker_name": workerName}
	err := client.connection.Send(internal.Request{
		RequestName: types.PruneWorker,
		Params:      params,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
	}, nil)

	if unexpectedResponseError, ok := err.(internal.UnexpectedResponseError); ok {
		if unexpectedResponseError.StatusCode == http.StatusBadRequest {
			var pruneWorkerErr PruneWorkerError

			err = json.Unmarshal([]byte(unexpectedResponseError.Body), &pruneWorkerErr)
			if err != nil {
				return err
			}

			return pruneWorkerErr
		}
	}

	return err
}

func (client *client) LandWorker(workerName string) error {
	params := rata.Params{"worker_name": workerName}
	err := client.connection.Send(internal.Request{
		RequestName: types.LandWorker,
		Params:      params,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
	}, nil)

	return err
}
