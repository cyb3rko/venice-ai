package characters_endpoints

import "net/http"

func List() *http.Request {
	path := "characters"
	req, _ := http.NewRequest("GET", path, nil)
	return req
}
