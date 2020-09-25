package concourse

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"io/ioutil"
	"net/http"

	"github.com/concourse/concourse/go-concourse/concourse/internal"
)

func (client *client) UserInfo() (types.UserInfo, error) {
	resp, err := client.httpAgent.Send(internal.Request{
		RequestName: types.GetUser,
	})

	if err != nil {
		return types.UserInfo{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var userInfo types.UserInfo
		err = json.NewDecoder(resp.Body).Decode(&userInfo)
		if err != nil {
			return types.UserInfo{}, err
		}
		return userInfo, nil
	case http.StatusUnauthorized:
		return types.UserInfo{}, ErrUnauthorized
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return types.UserInfo{}, internal.UnexpectedResponseError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}
}
