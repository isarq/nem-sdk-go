package requests

import (
	"fmt"
	"strings"
)

// An url encoded header
// type {object}
const UrlEncoded = `'Content-Type': 'application/x-www-form-urlencoded'`

// Create an application/json header
// param - A json string
// return - An application/json header with content length
func Json(data string) []string {
	slice := make([]string, 2)
	slice[0] = `"Content-Type": "application/json"`
	slice[1] = fmt.Sprintf("\"Content-Length\": %v", strings.NewReader(data).Len())
	return slice
}
