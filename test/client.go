package test

import (
	"os"
	"venice-ai/api"
)

var testClient *api.VeniceClient

func GetTestClient() *api.VeniceClient {
	if testClient == nil {
		testClient = api.NewClient(os.Getenv("VENICE_API_KEY"))
	}
	return testClient
}
