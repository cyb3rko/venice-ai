package api_keys_endpoints

import "net/http"

func RateLimits() *http.Request {
	path := "api_keys/rate_limits"
	req, _ := http.NewRequest("GET", path, nil)
	return req
}
