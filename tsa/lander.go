package tsa

import (
	"context"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"net/http/httputil"

	"fmt"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/tedsuo/rata"
)

type Lander struct {
	ATCEndpoint *rata.RequestGenerator
	HTTPClient  *http.Client
}

func (l *Lander) Land(ctx context.Context, worker types.Worker) error {
	logger := lagerctx.FromContext(ctx)

	logger.Info("start")
	defer logger.Info("end")

	request, err := l.ATCEndpoint.CreateRequest(types.LandWorker, rata.Params{
		"worker_name": worker.Name,
	}, nil)
	if err != nil {
		logger.Error("failed-to-construct-request", err)
		return err
	}

	response, err := l.HTTPClient.Do(request)
	if err != nil {
		logger.Error("failed-to-land", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logger.Error("bad-response", nil, lager.Data{
			"status-code": response.StatusCode,
		})

		b, _ := httputil.DumpResponse(response, true)
		return fmt.Errorf("bad-response (%d): %s", response.StatusCode, string(b))
	}

	return nil
}
