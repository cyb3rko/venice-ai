package tests

import (
	"testing"
	"venice-ai/images/images_endpoints"
	"venice-ai/test"
)

func TestImagesStyles(t *testing.T) {
	req := images_endpoints.Styles()
	resp, err := test.GetTestClient().Do(req)
	test.AssertHttp200(t, resp, err)
	test.AssertJson(t, resp)
}
