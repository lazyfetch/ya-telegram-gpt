package utils

import (
	"fmt"
	"main/telegram-bot/lang"
	"os"
)

func MergeID(id int64) string {
	idJSON := fmt.Sprintf("%d.json", id)
	return idJSON
}

func IsThereFile(id int64) bool {

	file, err := os.Open(fmt.Sprintf("./ai-models/yandexgpt/history/%s", MergeID(id)))
	if err != nil {
		defer file.Close()
		return false
	} else {
		defer file.Close()
		return true
	}

}

func CreateFile(id int64) {

	if !IsThereFile(id) {
		// test string
		// Copy base.json file
		file, _ := os.Open("./ai-models/yandexgpt/history/base.json")
		data, _ := os.ReadFile(file.Name())

		// Paste data
		newUserFile, _ := os.Create(fmt.Sprintf("./ai-models/yandexgpt/history/%s", MergeID(id)))
		newUserFile.Write(data)

		// Exit from file's
		defer file.Close()
		defer newUserFile.Close()
	}
}

func DeleteHistory(id int64) string {

	if !IsThereFile(id) {
		return lang.DELETE_CONTEXT_ERR_NOT_EXIST
	} else {
		_ = os.Remove(fmt.Sprintf("./ai-models/yandexgpt/history/%s", MergeID(id)))
		return lang.DELETE_CONTEXT_SUCCESFUL
	}

}
