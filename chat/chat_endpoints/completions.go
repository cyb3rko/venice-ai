package chat_endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"venice-ai/chat"
)

func Completion(config *chat.ChatConfig) *http.Request {
	path := "chat/completions"
	body, _ := json.Marshal(config)
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))
	return req
}
