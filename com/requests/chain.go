package requests

import (
	"encoding/json"
	"errors"
	"github.com/isarq/nem-sdk-go/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "github.com/isarq/nem-sdk-go/base"
)

// BlockHeight contains a chain height
type BlockHeight struct {
	Height int64
}

type TimeStamps struct {
	SendTimeStamp    int64 `json:"sendTimeStamp"`
	ReceiveTimeStamp int64 `json:"receiveTimeStamp"`
}

type Client struct {
	Node    Node
	URL     url.URL
	Request func(*http.Request) ([]byte, error)
}

type Block struct {
	TimeStamp     int64         `json:"timeStamp"`
	Signature     string        `json:"signature"`
	PrevBlockHash PrevBlockHash `json:"prevBlockHash"`
	Type          int           `json:"type"`
	Transactions  []interface{} `json:"transactions"`
	Version       int           `json:"version"`
	Signer        string        `json:"signer"`
	Height        int64         `json:"height"`
}

type PrevBlockHash struct {
	Data string `json:"data"`
}

type ExplorerBlockViewModel struct {
	Txes       []ExplorerTransferViewModel `json:"txes"`
	Block      Block                       `json:"block"`
	Hash       string                      `json:"hash"`
	Difficulty int                         `json:"difficulty"`
}

type ExplorerTransferViewModel struct {
	Tx        TransactionResponse `json:"tx"`
	Hash      string              `json:"hash"`
	InnerHash string              `json:"innerHash"`
}

func NewClient(node Node) *Client {
	protocol := strings.Split(node.Host, "://")
	host := utils.FormatEndpoint(node)

	return &Client{Node: node, URL: url.URL{Scheme: protocol[0], Host: host}}
}

// Gets the current height of the block chain.
// method Client - An Client endpoint struct point
// return {struct} - A [BlockHeight] struct
// link http://bob.nem.ninja/docs/#block-chain-height
func (c *Client) Height() (BlockHeight, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/chain/height"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return BlockHeight{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return BlockHeight{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return BlockHeight{}, err
	}

	var data BlockHeight
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return BlockHeight{}, err
	}
	return data, nil
}

// Gets the current last block of the chain.
// method Client - An Client endpoint struct point
// return
func (c *Client) LastBlock() (BlockHeight, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/chain/last-block"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return BlockHeight{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return BlockHeight{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return BlockHeight{}, err
	}

	var data BlockHeight
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return BlockHeight{}, err
	}
	return data, nil
}

// Gets network time (in ms)
// method Client - An Client endpoint struct point
// return - A [communicationTimeStamps]
// link http://bob.nem.ninja/docs/#communicationTimeStamps
func (c *Client) Time() (TimeStamps, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/time-sync/network-time"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return TimeStamps{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return TimeStamps{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return TimeStamps{}, err
	}

	var data TimeStamps
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return TimeStamps{}, err
	}
	return data, nil
}

// Gets a block by its height
// param Client - An Client endpoint struct point
// param height - The height of the block
// return - A block struct
func (c *Client) BlockByHeight(height int64) (Block, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	payload, err := json.Marshal(map[string]int64{"height": height})
	if err != nil {
		return Block{}, err
	}

	c.URL.Path = "/block/at/public"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return Block{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return Block{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return Block{}, err
	}

	var data Block
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return Block{}, err
	}
	return data, nil
}

// Gets part of a chain
// param Client - An Client endpoint struct point
// param height - The height of the block
// return - A array of ExplorerBlockViewModel struct
// link https://nemproject.github.io/#getting-part-of-a-chain
func (c *Client) BlockAfterByHeight(height int64) ([]ExplorerBlockViewModel, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	payload, err := json.Marshal(map[string]int64{"height": height})
	if err != nil {
		return []ExplorerBlockViewModel{}, err
	}

	c.URL.Path = "/local/chain/blocks-after"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []ExplorerBlockViewModel{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []ExplorerBlockViewModel{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []ExplorerBlockViewModel{}, err
	}

	var data struct {
		Datas []ExplorerBlockViewModel `json:"data"`
	}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return []ExplorerBlockViewModel{}, err
	}
	return data.Datas, nil
}
