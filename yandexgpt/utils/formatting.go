package utils

import (
	"main/yandexgpt/utils" .
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

func ReqFormatting(jsonID, prompt string, id int64) *bytes.Buffer {

	// Сheck if the file exists or create it
	CreateFile(id)

	file, _ := os.Open(fmt.Sprintf("./ai-models/yandexgpt/history/%s", jsonID))

	defer file.Close()

	// Read data from file
	dat, _ := os.ReadFile(file.Name())

	var data map[string]interface{}

	_ = json.Unmarshal(dat, &data)

	if messages, ok := data["messages"].([]interface{}); ok {
		// Create new message "user"
		newMessage := Message{
			Role: "user",
			Text: prompt,
		}

		// Convert object in map
		newMessageMap := map[string]interface{}{
			"role": newMessage.Role,
			"text": newMessage.Text,
		}

		// Append new message
		data["messages"] = append(messages, newMessageMap)
	}
	// Сериализуем обновленный JSON
	updatedData, _ := json.MarshalIndent(data, "", "  ")

	_ = os.WriteFile(fmt.Sprintf("./ai-models/yandexgpt/history/%s", jsonID), updatedData, 0644)

	// Reread for take actual data
	dat, _ = os.ReadFile(file.Name())

	return bytes.NewBuffer(dat)

}

// Due to the peculiarity of YandexGPT,
// we need to save our queries and responses to preserve the context of the dialog.
// This function implements this
func RespFormatting(jsonID, resBody string) {

	// Open and read
	file, _ := os.Open(fmt.Sprintf("./ai-models/yandexgpt/history/%s", jsonID))
	dat, _ := os.ReadFile(file.Name())
	defer file.Close()

	// Unmarshal
	var data map[string]interface{}
	_ = json.Unmarshal(dat, &data)

	if messages, ok := data["messages"].([]interface{}); ok {
		// Create new message "assistant"
		newMessage := Message{
			Role: "assistant",
			Text: resBody,
		}

		// Convert from object to map
		newMessageMap := map[string]interface{}{
			"role": newMessage.Role,
			"text": newMessage.Text,
		}

		// Append new message
		data["messages"] = append(messages, newMessageMap)
	}

	updatedData, _ := json.MarshalIndent(data, "", "  ")

	_ = os.WriteFile(fmt.Sprintf("./ai-models/yandexgpt/history/%s", jsonID), updatedData, 0644)

}
