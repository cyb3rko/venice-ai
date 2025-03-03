package chat_endpoints

import "net/http"

func Completion(model string) *http.Request {
	path := "api_keys/rate_limits"
	req, _ := http.NewRequest("POST", path, nil)
	return req
}
