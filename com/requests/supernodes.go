package requests

import (
	"encoding/json"
	"errors"
	"github.com/isarq/nem-sdk-go/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SuperNodeInfo struct {
	Nodes     []Supernode `json:"nodes"`
	NodeCount int         `json:"nodeCount"`
}

type Supernode struct {
	ID            string  `json:"id"`
	Alias         string  `json:"alias"`
	IP            string  `json:"ip"`
	NisPort       int     `json:"nisPort"`
	PubKey        string  `json:"pubKey"`
	ServantPort   int     `json:"servantPort"`
	Status        int     `json:"status"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	PayoutAddress string  `json:"payoutAddress"`
}

type SuperNode struct {
	ID            int     `json:"id,omitempty"`
	Alias         string  `json:"alias,omitempty"`
	IP            string  `json:"ip,omitempty"`
	NisPort       int     `json:"nisPort,omitempty"`
	PubKey        string  `json:"pubKey,omitempty"`
	ServantPort   int     `json:"servantPort,omitempty"`
	Status        int     `json:"status,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	PayoutAddress string  `json:"payoutAddress,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Distance      float64 `json:"distance,omitempty"`
	MaxUnlocked   int     `json:"maxUnlocked,omitempty"`
	NumUnlocked   int     `json:"numUnlocked,omitempty"`
}

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	NumNodes  int     `json:"numNodes"`
}

// Gets all nodes of the node reward program
// return - An SuperNodeInfo struct
func SuperNodeAll() (SuperNodeInfo, error) {
	c := Client{}
	timeout := time.Duration(50 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	node := strings.Split(model.Supernodes, "//")
	node = strings.Split(node[1], "/")
	c.URL.Host = node[0]
	c.URL.Path = "/nodes"
	c.URL.Scheme = "https"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return SuperNodeInfo{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return SuperNodeInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return SuperNodeInfo{}, err
	}

	var data = SuperNodeInfo{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return SuperNodeInfo{}, err
	}
	return data, nil
}

// Gets the nearest supernodes
// param coords - A coordinates object: https://www.w3schools.com/html/html5_geolocation.asp
// return - An SuperNodeInfo struct
func Nearest(latitude, longitude float64) ([]SuperNode, error) {
	c := Client{}
	timeout := time.Duration(20 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	Ojt := Coords{Latitude: latitude, Longitude: longitude, NumNodes: 5}

	payload, err := json.Marshal(Ojt)
	if err != nil {
		return []SuperNode{}, err
	}
	c.URL.Host = "199.217.113.179:7782"
	c.URL.Path = "/nodes/nearest"
	c.URL.Scheme = "http"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []SuperNode{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []SuperNode{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []SuperNode{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct{ Data []SuperNode }{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return []SuperNode{}, err
	}
	return data.Data, nil
}

// Gets the all supernodes by status
// param {number} status - 0 for all nodes, 1 for active nodes, 2 for inactive nodes
// return {struct} - An SuperNodeInfo struct
func GetSuperNodeStatus(status int) ([]SuperNode, error) {
	c := Client{}
	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if status > 2 {
		err := errors.New("error: status - 0 for all nodes, 1 for active nodes, 2 for inactive nodes")
		return []SuperNode{}, err
	}

	payload, err := json.Marshal(map[string]int{"status": status})
	if err != nil {
		return []SuperNode{}, err
	}
	c.URL.Host = "199.217.113.179:7782"
	c.URL.Path = "/nodes"
	c.URL.Scheme = "http"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []SuperNode{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []SuperNode{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []SuperNode{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct {
		Data []SuperNode
	}{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return []SuperNode{}, err
	}
	return data.Data, nil
}
