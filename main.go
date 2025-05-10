package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"venice-ai/api"
	"venice-ai/chat"
	"venice-ai/chat/chat_endpoints"
	"venice-ai/images"
	"venice-ai/images/images_endpoints"
	"venice-ai/utils"
)

func main() {
	client := api.NewClient(os.Getenv("VENICE_API_KEY"))

	completeChat(client)
	//generateImage(client)
}

func completeChat(client *api.VeniceClient) {
	req := chat_endpoints.Completion(&chat.ChatConfig{
		Model: "mistral-31-24b",
		Messages: []map[string]interface{}{
			{
				"content": "Hello, who are you?",
				"role":    "user",
			},
		},
		Stream: false,
	})
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	utils.PrettyPrint(body)
}

func generateImage(client *api.VeniceClient) {
	res, err := client.Do(images_endpoints.Generate(&images.ImageConfig{
		Model:  "fluently-xl",
		Prompt: "Cute dog",
	}))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var payload images_endpoints.GenerateResponse
	err = json.Unmarshal(body, &payload)
	_ = utils.PersistBase64ToFile(payload.Id, payload.Request.Format, payload.Images[0])
	for i := range payload.Images {
		payload.Images[i] = "..."
	}
	utils.PrettyPrintIf(payload)
}
