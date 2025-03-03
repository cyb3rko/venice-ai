package images_endpoints

import "net/http"

func Styles() *http.Request {
	path := "image/styles"
	req, _ := http.NewRequest("GET", path, nil)
	return req
}
