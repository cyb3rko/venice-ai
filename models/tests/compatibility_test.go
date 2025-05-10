package tests

import (
	"testing"
	"venice-ai/models"
	"venice-ai/models/models_endpoints"
	"venice-ai/test"
)

func TestModelsCompatibilityMapping(t *testing.T) {
	req := models_endpoints.CompatibilityMapping(models.ModelTypeDefault)
	resp, err := test.GetTestClient().Do(req)
	test.AssertHttp200(t, resp, err)
	test.AssertJson(t, resp)
}
