package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type NemRequestResult struct {
	Type    int    `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Determines if NIS is up and responsive.
// method Client - An Client endpoint struct point
// return - A [NemRequestResult] struct
// link http://bob.nem.ninja/docs/#nemRequestResult
func (c *Client) Heartbeat() (NemRequestResult, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/heartbeat"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return NemRequestResult{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return NemRequestResult{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return NemRequestResult{}, err
	}

	var data NemRequestResult
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return NemRequestResult{}, err
	}
	return data, nil
}
