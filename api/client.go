package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type VeniceClient struct {
	httpClient *http.Client
	baseUrl    string
	headers    map[string]string
}

func NewClient(apiKey string) *VeniceClient {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", apiKey),
		"Content-Type":  "application/json",
	}
	return &VeniceClient{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		headers:    headers,
	}
}

func (client *VeniceClient) Do(req *http.Request) (*http.Response, error) {
	fullUrl, err := url.Parse(client.baseUrl + req.URL.String())
	if err != nil {
		return nil, err
	}
	req.URL = fullUrl
	for key, value := range client.headers {
		req.Header.Set(key, value)
	}
	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, err
}
