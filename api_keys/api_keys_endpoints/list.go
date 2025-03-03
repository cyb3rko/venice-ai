package api_keys_endpoints

import (
	"net/http"
)

func List() *http.Request {
	path := "api_keys"
	req, _ := http.NewRequest("GET", path, nil)
	return req
}
