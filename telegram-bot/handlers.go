package telegrambot

import (
	"main/telegram-bot/lang"
	ya "main/yandexgpt"
	"main/yandexgpt/utils"

	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func DeleteContext(bh *th.BotHandler) {

	bh.Handle(func(bot *tg.Bot, update tg.Update) {
		_, _ = bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), utils.DeleteHistory(update.Message.Chat.ID),
		))
	}, th.CommandEqual("deletecontext"))

}

// start command handler
func StartCommandHandler(bh *th.BotHandler) {

	bh.Handle(func(bot *tg.Bot, update tg.Update) {
		_, _ = bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), lang.START,
		))
	}, th.CommandEqual("start"))
}

func PromptHandler(bh *th.BotHandler) {
	bh.HandleMessage(func(bot *tg.Bot, msg tg.Message) {
		_, _ = bot.SendMessage(tu.Message(msg.Chat.ChatID(), ya.Requests(msg.Text, msg.Chat.ChatID().ID)))
	})
}
