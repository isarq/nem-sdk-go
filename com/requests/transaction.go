package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// The NemAnnounceResult extends the NemRequestResult by supplying
// the additional fields 'transactionHash' and in case of a multisig transaction 'innerTransactionHash'.
type NemAnnounceResult struct {
	Type                 int                  `json:"type"`
	Code                 int                  `json:"code"`
	Message              string               `json:"message"`
	TransactionHash      TransactionHash      `json:"transactionHash,omitempty"`
	InnerTransactionHash InnerTransactionHash `json:"innerTransactionHash,omitempty"`
}

type TransactionHash struct {
	Data string `json:"data,omitempty"`
}

type InnerTransactionHash struct {
	Data string `json:"data,omitempty"`
}

// A RequestAnnounce object is used to transfer the transaction data and
// the signature to NIS in order to initiate and broadcast a transaction.
type RequestAnnounce struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

// Broadcast a transaction to the NEM network
// Method Client - An Client endpoint struct point
// param serialize - A RequestAnnounce struct
// return - A [NemAnnounceResult] struct
// link http://bob.nem.ninja/docs/#nemAnnounceResult
func (c Client) Announce(serialize RequestAnnounce) (NemAnnounceResult, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	payload, err := json.Marshal(serialize)
	if err != nil {
		return NemAnnounceResult{}, err
	}
	c.URL.Path = "/transaction/announce"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return NemAnnounceResult{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprintf("%v", len(payload)))

	resp, err := client.Do(req)
	if err != nil {
		return NemAnnounceResult{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return NemAnnounceResult{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := NemAnnounceResult{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return NemAnnounceResult{}, err
	}
	return data, nil
}

// Gets a TransactionMetaDataPair object from the chain using it's hash
// Method Client - An Client endpoint struct point
// param txHash - A transaction hash
// return A [TransactionMetaDataPair] struct
// link http://bob.nem.ninja/docs/#transactionMetaDataPair
func (c *Client) ByHash(txHash string) (base.Transaction, error) {
	b := new(bytes.Buffer)
	timeout := 10 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/transaction/get"
	req, err := c.buildReq(map[string]string{"hash": txHash}, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b := &bytes.Buffer{}
		_, _ = b.ReadFrom(resp.Body)
		return nil, errors.New(b.String())
	}

	_, _ = io.Copy(b, resp.Body)

	return MapTransaction(b)
}
