package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"io/ioutil"
	"net/http"
	"time"
)

type ID struct {
	NamespaceID string `json:"namespaceId,omitempty"`
	Name        string `json:"name,omitempty"`
}

type MosaicSupplyInfo struct {
	MosaicID base.MosaicID `json:"mosaicId"`
	Supply   int           `json:"supply"`
}

// A mosaic definition consists of a database id and a mosaic definition object.
// The id is needed for requests that support paging.
type MosaicDefinitionMetaDataPair struct {
	Meta   Meta                  `json:"meta"`
	Mosaic base.MosaicDefinition `json:"mosaic"`
}

// Gets the current supply of a mosaic
// method Client - An Client endpoint struct point
// param id - A mosaic id
// return - An mosaicSupplyInfo struct
func (c *Client) Supply(id string) (MosaicSupplyInfo, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/mosaic/supply"
	req, err := c.buildReq(map[string]string{"mosaicId": id}, nil, http.MethodGet)
	if err != nil {
		return MosaicSupplyInfo{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return MosaicSupplyInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return MosaicSupplyInfo{}, err
	}

	var data MosaicSupplyInfo
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return MosaicSupplyInfo{}, err
	}
	return data, nil
}
