package action

import (
  // "fmt"
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

func Run (bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)

    var msg tgbotapi.Chattable

    switch update.Message.Text {
      case "/help":
        msg = tgbotapi.NewMessage(update.Message.Chat.ID, "описание о боте")
      case "/settings":
        msg = tgbotapi.NewMessage(update.Message.Chat.ID, "описание настроек")
      default:
        msg = tgbotapi.NewMessage(update.Message.Chat.ID, "команда непонятна")
    }
    bot.Send(msg)
	}
}
