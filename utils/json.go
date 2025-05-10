package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func PrettyPrint(data []byte) {
	out := &bytes.Buffer{}
	if err := json.Indent(out, data, "", "  "); err == nil {
		fmt.Println(out.String())
	} else {
		fmt.Println(string(data))
	}
}

func PrettyPrintIf(data interface{}) {
	dataBytes, _ := json.Marshal(data)
	PrettyPrint(dataBytes)
}

func PrettyPrintTest(t *testing.T, data []byte) {
	out := &bytes.Buffer{}
	if err := json.Indent(out, data, "", "  "); err == nil {
		t.Error(out.String())
	} else {
		t.Error(string(data))
	}
}
