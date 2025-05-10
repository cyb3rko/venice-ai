package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"venice-ai/utils"
)

func AssertJson(t *testing.T, resp *http.Response) {
	var body map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Error("Response body is not valid JSON:")
		bodyString, _ := io.ReadAll(resp.Body)
		utils.PrettyPrintTest(t, bodyString)
	}
}

func AssertHttp200(t *testing.T, resp *http.Response, err error) {
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: expected %d, got %d", http.StatusOK, resp.StatusCode)
		t.Errorf("Response: %s", resp.Status)
	}
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}
