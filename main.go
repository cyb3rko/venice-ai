package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"venice-ai/images"
	"venice-ai/images/images_endpoints"
	"venice-ai/utils"
)

func main() {
	client := NewClient(os.Getenv("VENICE_API_KEY"))

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
	_ = utils.PersistBase64ToFile(payload.Id, payload.Images[0])
	for i := range payload.Images {
		payload.Images[i] = "..."
	}
	utils.PrettyPrint(payload)
}
