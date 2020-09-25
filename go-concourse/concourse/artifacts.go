package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) CreateArtifact(src io.Reader, platform string, tags []string) (types.WorkerArtifact, error) {
	var artifact types.WorkerArtifact

	params := rata.Params{
		"team_name": team.Name(),
	}

	err := team.connection.Send(internal.Request{
		Header:      http.Header{"Content-Type": {"application/octet-stream"}},
		RequestName: types.CreateArtifact,
		Params:      params,
		Query:       url.Values{"platform": {platform}, "tags": tags},
		Body:        src,
	}, &internal.Response{
		Result: &artifact,
	})

	return artifact, err
}

func (team *team) GetArtifact(artifactID int) (io.ReadCloser, error) {
	params := rata.Params{
		"team_name":   team.Name(),
		"artifact_id": strconv.Itoa(artifactID),
	}

	response := internal.Response{}
	err := team.connection.Send(internal.Request{
		RequestName:        types.GetArtifact,
		Params:             params,
		ReturnResponseBody: true,
	}, &response)

	if err != nil {
		return nil, err
	}

	return response.Result.(io.ReadCloser), nil
}
