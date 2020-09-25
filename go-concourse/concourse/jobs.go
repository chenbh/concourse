package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"net/http"
	"net/url"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) ListJobs(pipelineName string) ([]types.Job, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"team_name":     team.name,
	}

	var jobs []types.Job
	err := team.connection.Send(internal.Request{
		RequestName: types.ListJobs,
		Params:      params,
	}, &internal.Response{
		Result: &jobs,
	})

	return jobs, err
}

func (client *client) ListAllJobs() ([]types.Job, error) {
	var jobs []types.Job
	err := client.connection.Send(internal.Request{
		RequestName: types.ListAllJobs,
	}, &internal.Response{
		Result: &jobs,
	})

	return jobs, err
}

func (team *team) Job(pipelineName, jobName string) (types.Job, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	var job types.Job
	err := team.connection.Send(internal.Request{
		RequestName: types.GetJob,
		Params:      params,
	}, &internal.Response{
		Result: &job,
	})
	switch err.(type) {
	case nil:
		return job, true, nil
	case internal.ResourceNotFoundError:
		return job, false, nil
	default:
		return job, false, err
	}
}

func (team *team) JobBuilds(pipelineName string, jobName string, page Page) ([]types.Build, Pagination, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	var builds []types.Build

	headers := http.Header{}
	err := team.connection.Send(internal.Request{
		RequestName: types.ListJobBuilds,
		Params:      params,
		Query:       page.QueryParams(),
	}, &internal.Response{
		Result:  &builds,
		Headers: &headers,
	})
	switch err.(type) {
	case nil:
		pagination, err := paginationFromHeaders(headers)
		if err != nil {
			return builds, Pagination{}, false, err
		}

		return builds, pagination, true, nil
	case internal.ResourceNotFoundError:
		return builds, Pagination{}, false, nil
	default:
		return builds, Pagination{}, false, err
	}
}

func (team *team) PauseJob(pipelineName string, jobName string) (bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: types.PauseJob,
		Params:      params,
	}, &internal.Response{})

	switch err.(type) {
	case nil:
		return true, nil
	case internal.ResourceNotFoundError:
		return false, nil
	default:
		return false, err
	}
}

func (team *team) UnpauseJob(pipelineName string, jobName string) (bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: types.UnpauseJob,
		Params:      params,
	}, &internal.Response{})

	switch err.(type) {
	case nil:
		return true, nil
	case internal.ResourceNotFoundError:
		return false, nil
	default:
		return false, err
	}
}

func (team *team) ScheduleJob(pipelineName string, jobName string) (bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"team_name":     team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: types.ScheduleJob,
		Params:      params,
	}, &internal.Response{})

	switch err.(type) {
	case nil:
		return true, nil
	case internal.ResourceNotFoundError:
		return false, nil
	default:
		return false, err
	}
}

func (team *team) ClearTaskCache(pipelineName string, jobName string, stepName string, cachePath string) (int64, error) {
	params := rata.Params{
		"team_name":     team.name,
		"pipeline_name": pipelineName,
		"job_name":      jobName,
		"step_name":     stepName,
	}

	queryParams := url.Values{}
	if len(cachePath) > 0 {
		queryParams.Add(types.ClearTaskCacheQueryPath, cachePath)
	}

	var ctcResponse types.ClearTaskCacheResponse
	responseHeaders := http.Header{}
	response := internal.Response{
		Headers: &responseHeaders,
		Result:  &ctcResponse,
	}
	err := team.connection.Send(internal.Request{
		RequestName: types.ClearTaskCache,
		Params:      params,
		Query:       queryParams,
	}, &response)

	if err != nil {
		return 0, err
	} else {
		return ctcResponse.CachesRemoved, nil
	}
}
