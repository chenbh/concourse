package concourse

import (
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (client *client) Check(checkID string) (types.Check, bool, error) {

	params := rata.Params{
		"check_id": checkID,
	}

	var check types.Check
	err := client.connection.Send(internal.Request{
		RequestName: types.GetCheck,
		Params:      params,
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
