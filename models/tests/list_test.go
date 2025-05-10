package tests

import (
	"testing"
	"venice-ai/models"
	"venice-ai/models/models_endpoints"
	"venice-ai/test"
)

func TestModelsList(t *testing.T) {
	req := models_endpoints.List(models.ModelTypeDefault)
	resp, err := test.GetTestClient().Do(req)
	test.AssertHttp200(t, resp, err)
	test.AssertJson(t, resp)
}
