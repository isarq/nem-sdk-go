package requests

import (
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Audit an apostille file
// param publicKey - The signer public key
// param data - The file data of audited file
// param signedData - The signed data into the apostille transaction message
// return - True if valid, false otherwise
func Audit(publicKey, data, signedData string) (bool, error) {
	c := Client{}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	node := strings.Split(model.ApostilleAuditServer, "//")
	node = strings.Split(node[1], "/")
	c.URL.Host = node[0]
	c.URL.Path = "/verify"
	c.URL.Scheme = "http"
	params := make(map[string]string)
	params["publicKey"] = publicKey
	params["data"] = data
	params["signedData"] = signedData
	req, err := c.buildReq(params, nil, http.MethodPost)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(byteArray))

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return false, err
	}

	return true, nil
}
