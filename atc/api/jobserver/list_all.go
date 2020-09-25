package jobserver

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/api/present"
)

func (s *Server) ListAllJobs(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("list-all-jobs")

	acc := accessor.GetAccessor(r)

	var dashboard atc.Dashboard
	var err error

	if acc.IsAdmin() {
		dashboard, err = s.jobFactory.AllActiveJobs()
	} else {
		dashboard, err = s.jobFactory.VisibleJobs(acc.TeamNames())
	}

	if err != nil {
		logger.Error("failed-to-get-all-visible-jobs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jobs := []types.Job{}

	for _, job := range dashboard {
		jobs = append(
			jobs,
			present.DashboardJob(
				job.TeamName,
				job,
			),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(jobs)
	if err != nil {
		logger.Error("failed-to-encode-jobs", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
