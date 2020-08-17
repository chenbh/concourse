package buildserver

import (
	"net/http"

	"github.com/chenbh/concourse/v6/atc/db"
)

func (s *Server) BuildEvents(build db.Build) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		streamDone := make(chan struct{})

		go func() {
			defer close(streamDone)

			s.eventHandlerFactory(s.logger, build).ServeHTTP(w, r)
		}()

		select {
		case <-streamDone:
		}
	})
}
