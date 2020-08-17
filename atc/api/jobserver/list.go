package jobserver

import (
	"encoding/json"
	"net/http"

	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/api/present"
	"github.com/chenbh/concourse/v6/atc/db"
)

func (s *Server) ListJobs(pipeline db.Pipeline) http.Handler {
	logger := s.logger.Session("list-jobs")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jobs := []atc.Job{}

		dashboard, err := pipeline.Dashboard()

		if err != nil {
			logger.Error("failed-to-get-dashboard", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		teamName := r.FormValue(":team_name")

		for _, job := range dashboard {
			jobs = append(
				jobs,
				present.DashboardJob(
					teamName,
					job,
				),
			)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(jobs)
		if err != nil {
			logger.Error("failed-to-encode-jobs", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
