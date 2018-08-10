package requests

import (
	"bytes"
	"net/http"
)

func (c Client) buildReq(params map[string]string, body []byte, method string) (*http.Request, error) {
	if params != nil {
		q := c.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		c.URL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, c.URL.String(), bytes.NewBuffer(body))
	if err != nil {
		return &http.Request{}, err
	}
	return req, nil
}
