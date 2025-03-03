package models_endpoints

import (
	"net/http"
	"net/url"
	"venice-ai/models"
)

func CompatibilityMapping(modelType models.ModelType) *http.Request {
	path := "models/compatibility_mapping"
	req, _ := http.NewRequest("GET", path, nil)
	query := url.Values{}
	query.Set("type", string(modelType))
	req.URL.RawQuery = query.Encode()
	return req
}
