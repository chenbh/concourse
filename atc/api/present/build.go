package present

import (
	"github.com/concourse/concourse/atc/types"
	"strconv"

	"github.com/concourse/concourse/atc/db"
	"github.com/tedsuo/rata"
)

func Build(build db.Build) types.Build {

	apiURL, err := types.Routes.CreatePathForRoute(types.GetBuild, rata.Params{
		"build_id":  strconv.Itoa(build.ID()),
		"team_name": build.TeamName(),
	})
	if err != nil {
		panic("failed to generate url: " + err.Error())
	}

	atcBuild := types.Build{
		ID:           build.ID(),
		Name:         build.Name(),
		JobName:      build.JobName(),
		PipelineName: build.PipelineName(),
		TeamName:     build.TeamName(),
		Status:       string(build.Status()),
		APIURL:       apiURL,
	}

	if build.RerunOf() != 0 {
		atcBuild.RerunNumber = build.RerunNumber()
		atcBuild.RerunOf = &types.RerunOfBuild{
			Name: build.RerunOfName(),
			ID:   build.RerunOf(),
		}
	}

	if !build.StartTime().IsZero() {
		atcBuild.StartTime = build.StartTime().Unix()
	}

	if !build.EndTime().IsZero() {
		atcBuild.EndTime = build.EndTime().Unix()
	}

	if !build.ReapTime().IsZero() {
		atcBuild.ReapTime = build.ReapTime().Unix()
	}

	return atcBuild
}
