package wrappa

import (
	"code.cloudfoundry.org/lager"
	"github.com/NYTimes/gziphandler"
	"github.com/concourse/concourse/atc/types"
	"github.com/tedsuo/rata"
)

type CompressionWrappa struct {
	lager.Logger
}

func NewCompressionWrappa(logger lager.Logger) Wrappa {
	return CompressionWrappa{
		logger,
	}
}

func (wrappa CompressionWrappa) Wrap(handlers rata.Handlers) rata.Handlers {
	wrapped := rata.Handlers{}

	for name, handler := range handlers {
		switch name {
		// always gzip for events
		case types.BuildEvents:
			gzipEnforcedHandler, err := gziphandler.GzipHandlerWithOpts(gziphandler.MinSize(0))
			if err != nil {
				wrappa.Logger.Error("failed-to-create-gzip-handler", err)
			}

			wrapped[name] = gzipEnforcedHandler(handler)
		// skip gzip as this endpoint does it already
		case types.DownloadCLI:
			wrapped[name] = handler
		default:
			wrapped[name] = gziphandler.GzipHandler(handler)
		}
	}

	return wrapped
}
