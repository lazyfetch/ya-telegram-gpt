package main

import (
	"fmt"
	"os"

	telegrambot "main/telegram-bot"

	goenv "github.com/joho/godotenv"
	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {

	env, _ := goenv.Read("config.env")
	API_KEY := env["TELEGRAM_APIKEY"]

	// create bot
	bot, err := tg.NewBot(API_KEY, tg.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// use Long Polling for take response
	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	// handler's

	telegrambot.DeleteContext(bh)
	telegrambot.StartCommandHandler(bh)
	telegrambot.PromptHandler(bh)

	bh.Start()

}
