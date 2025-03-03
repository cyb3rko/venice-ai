package utils

import (
	"bytes"
	"encoding/json"
)

func PrettyPrint(data interface{}) {
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err == nil {
		println(buffer.String())
	} else {
		println(data)
	}
}
