package infoserver

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"net/http"
)

func (s *Server) Info(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("info")

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(types.Info{Version: s.version,
		WorkerVersion: s.workerVersion,
		ExternalURL:   s.externalURL,
		ClusterName:   s.clusterName,
	})
	if err != nil {
		logger.Error("failed-to-encode-info", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
