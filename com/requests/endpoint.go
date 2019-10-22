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

type NemNodeInfo struct {
	MetaData struct {
		Features    int         `json:"features"`
		Application interface{} `json:"application"`
		NetworkID   int         `json:"networkId"`
		Version     string      `json:"version"`
		Platform    string      `json:"platform"`
	} `json:"metaData"`
	Endpoint struct {
		Protocol string `json:"protocol"`
		Port     int    `json:"port"`
		Host     string `json:"host"`
	} `json:"endpoint"`
	Identity struct {
		Name      string `json:"name"`
		PublicKey string `json:"public-key"`
	} `json:"identity"`
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

func (c *Client) GetNodeInfo() (NemNodeInfo, error) {
	timeout := 10 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/node/info"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return NemNodeInfo{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return NemNodeInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return NemNodeInfo{}, err
	}

	var data NemNodeInfo
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return NemNodeInfo{}, err
	}
	return data, nil
}
