package images_endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"venice-ai/images"
)

type GenerateResponse struct {
	Id      string             `json:"id"`
	Request images.ImageConfig `json:"request"`
	Images  []string           `json:"images"`
}

func Generate(config *images.ImageConfig) *http.Request {
	path := "image/generate"
	body, _ := json.Marshal(config)
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))
	return req
}
