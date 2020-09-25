package wrappa

import (
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/metric"
	"github.com/concourse/concourse/atc/types"
	"github.com/tedsuo/rata"
)

type APIMetricsWrappa struct {
	logger lager.Logger
}

func NewAPIMetricsWrappa(logger lager.Logger) Wrappa {
	return APIMetricsWrappa{
		logger: logger,
	}
}

func (wrappa APIMetricsWrappa) Wrap(handlers rata.Handlers) rata.Handlers {
	wrapped := rata.Handlers{}

	for name, handler := range handlers {
		switch name {
		case types.BuildEvents, types.DownloadCLI, types.HijackContainer:
			wrapped[name] = handler
		default:
			wrapped[name] = metric.WrapHandler(
				wrappa.logger,
				metric.Metrics,
				name,
				handler,
			)
		}
	}

	return wrapped
}
