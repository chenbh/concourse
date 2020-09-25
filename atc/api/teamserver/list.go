package teamserver

import (
	"encoding/json"
	"errors"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/api/present"
)

func (s *Server) ListTeams(w http.ResponseWriter, r *http.Request) {
	hLog := s.logger.Session("list-teams")

	teams, err := s.teamFactory.GetTeams()
	if err != nil {
		hLog.Error("failed-to-get-teams", errors.New("sorry"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	acc := accessor.GetAccessor(r)
	presentedTeams := make([]types.Team, 0)
	for _, team := range teams {
		if acc.IsAuthorized(team.Name()) {
			presentedTeams = append(presentedTeams, present.Team(team))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(presentedTeams)
	if err != nil {
		hLog.Error("failed-to-encode-teams", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
