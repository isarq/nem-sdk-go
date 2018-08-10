package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// A namespace consists of a namespace object and a database id.
// the id is needed for requests that support paging.
type NamespaceMetaDataPair struct {
	Meta      Meta      `json:"meta"`
	Namespace Namespace `json:"namespace"`
}

var algo base.MosaicDefinition

type Meta struct {
	ID int `json:"id"`
}

// A namespace consists of a namespace object and a database id. the id is needed for
// requests that support paging.
type Namespace struct {
	Fqn    string `json:"fqn"`
	Owner  string `json:"owner"`
	Height int    `json:"height"`
}

// Gets root namespaces.
// method Client - An Client endpoint struct point
// param id - The namespace id up to which root namespaces are returned (optional)
// return - An slice of [NamespaceMetaDataPair] struct
// link http://bob.nem.ninja/docs/#namespaceMetaDataPair
func (c *Client) NameSpaceRoots(id int) ([]NamespaceMetaDataPair, error) {

	params := map[string]string{"pageSize": "100"}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if id != 0 {
		Id := strconv.Itoa(id)
		params["id"] = Id
	}

	c.URL.Path = "/namespace/root/page"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []NamespaceMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []NamespaceMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []NamespaceMetaDataPair{}, err
	}

	var data = struct{ Data []NamespaceMetaDataPair }{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []NamespaceMetaDataPair{}, err
	}
	return data.Data, nil
}

// Gets mosaic definitions of a namespace
// method Client - An Client endpoint struct point
// param id - A namespace id
// return - An slice of [MosaicDefinition] struct
// link http://bob.nem.ninja/docs/#mosaicDefinition
func (c *Client) MosaicDefinitions(id string) ([]MosaicDefinitionMetaDataPair, error) {
	params := map[string]string{"namespace": id}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/namespace/mosaic/definition/page"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []MosaicDefinitionMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []MosaicDefinitionMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []MosaicDefinitionMetaDataPair{}, err
	}

	var data = struct {
		Data []MosaicDefinitionMetaDataPair
	}{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []MosaicDefinitionMetaDataPair{}, err
	}
	if data.Data == nil {
		return []MosaicDefinitionMetaDataPair{}, err
	}

	return data.Data, nil
}

// Gets the namespace with given id.
// method Client - An Client endpoint struct point
// param id - A namespace id
// return - A [NamespaceInfo] struct
// link http://bob.nem.ninja/docs/#namespace
func (c *Client) Namespaceinfo(id string) (Namespace, error) {
	params := map[string]string{"namespace": id}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/namespace"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return Namespace{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Namespace{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return Namespace{}, err
	}

	var data Namespace
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return Namespace{}, err
	}
	return data, nil
}

// Search for mosaic definition(s) into an array of mosaicDefinition objects
// param mosaicDefinitionArray - An array of mosaicDefinition objects
// param keys - Array of strings with names of the mosaics to find (['eur', 'usd',...])
// return - An object of mosaicDefinition objects
func SearchMosaicDefinitionArray(mosaicDefinitionArray []MosaicDefinitionMetaDataPair,
	keys []string) map[string]base.MosaicDefinition {
	rest := make(map[string]base.MosaicDefinition)
	for _, b := range keys {
		for _, c := range mosaicDefinitionArray {
			if c.Mosaic.ID.Name == b {
				m := utils.MosaicIdToName(c.Mosaic.ID)
				rest[m] = c.Mosaic
			}
		}
	}
	return rest
}
