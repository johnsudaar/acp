package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	apiResultSuccess = "Success"
)

type apiRequest struct {
	Request request `json:"Request"`
}

type request struct {
	Command   string      `json:"Command"`
	SessionID string      `json:"SessionID"`
	Params    interface{} `json:"Params"`
}

type apiResponse struct {
	Response response `json:"Response"`
}

type response struct {
	Requested string          `json:"Requested"`
	Result    string          `json:"Result"`
	Data      json.RawMessage `json:"Data"`
}

func (c HTTPClient) makeRequest(command string, params interface{}) (*response, error) {
	body, err := json.Marshal(apiRequest{
		Request: request{
			Command:   command,
			SessionID: c.sessionID,
			Params:    params,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "fail to build request body")
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/cgi-bin/api.cgi", c.ip), "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "fail to make HTTP request")
	}
	defer resp.Body.Close()

	var response apiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "invalid body")
	}

	if response.Response.Result != apiResultSuccess {
		return nil, fmt.Errorf("%s: Operation failed (%s)", c.ip, response.Response.Result)
	}

	return &response.Response, nil
}
