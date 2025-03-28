package yandexgpt

import (
	"encoding/json"
	"fmt"
	"io"
	"main/yandexgpt/utils"
	"net/http"

	goenv "github.com/joho/godotenv"
)

func Requests(prompt string, id int64) string {
	// test comm
	env, _ := goenv.Read("config.env")
	API_KEY := fmt.Sprintf("Api-key %s", env["YANDEX_APIKEY"])

	url := "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"

	req, err := http.NewRequest("POST", url, utils.ReqFormatting(utils.MergeID(id), prompt, id)) // ???
	req.Header.Set("Authorization", API_KEY)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data map[string]interface{}

	body, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(body, &data)

	v := data["result"].(map[string]interface{})["alternatives"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["text"].(string)
	utils.RespFormatting(utils.MergeID(id), v)
	return v

}
