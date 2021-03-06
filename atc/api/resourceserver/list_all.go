package resourceserver

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/api/present"
	"github.com/concourse/concourse/atc/db"
)

func (s *Server) ListAllResources(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("list-all-resources")

	acc := accessor.GetAccessor(r)

	var (
		dbResources []db.Resource
		err         error
	)
	if acc.IsAdmin() {
		dbResources, err = s.resourceFactory.AllResources()
	} else {
		dbResources, err = s.resourceFactory.VisibleResources(acc.TeamNames())
	}
	if err != nil {
		logger.Error("failed-to-get-all-visible-resources", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resources := []types.Resource{}

	for _, resource := range dbResources {
		resources = append(
			resources,
			present.Resource(
				resource,
				true,
				resource.TeamName(),
			),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resources)
	if err != nil {
		logger.Error("failed-to-encode-resources", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
