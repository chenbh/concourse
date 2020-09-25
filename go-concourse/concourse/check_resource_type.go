package concourse

import (
	"bytes"
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) CheckResourceType(pipelineName string, resourceTypeName string, version types.Version) (types.Check, bool, error) {

	params := rata.Params{
		"pipeline_name":      pipelineName,
		"resource_type_name": resourceTypeName,
		"team_name":          team.name,
	}

	var check types.Check

	jsonBytes, err := json.Marshal(types.CheckRequestBody{From: version})
	if err != nil {
		return check, false, err
	}

	err = team.connection.Send(internal.Request{
		RequestName: types.CheckResourceType,
		Params:      params,
		Body:        bytes.NewBuffer(jsonBytes),
		Header:      http.Header{"Content-Type": []string{"application/json"}},
	}, &internal.Response{
		Result: &check,
	})

	switch e := err.(type) {
	case nil:
		return check, true, nil
	case internal.ResourceNotFoundError:
		return check, false, nil
	case internal.UnexpectedResponseError:
		if e.StatusCode == http.StatusInternalServerError {
			return check, false, GenericError{e.Body}
		} else {
			return check, false, err
		}
	default:
		return check, false, err
	}
}
