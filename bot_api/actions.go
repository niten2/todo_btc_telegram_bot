package bot_api

import (
  "fmt"
  "log"
  "gopkg.in/telegram-bot-api.v4"
  "regexp"
)

func InitActions (bot *tgbotapi.BotAPI) {
  fmt.Println("InitActions for bot")

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

    user_name := update.Message.From.UserName
    chat_id := update.Message.Chat.ID
    text := update.Message.Text

    log.Printf("[%s] %s", user_name, text, chat_id)

    message := CreateResponse(text)
    response := tgbotapi.NewMessage(chat_id, message)

    bot.Send(response)
  }
}

func CreateResponse(input string) string {
    var msg string

    res, _ := regexp.MatchString("^p ", input)

    if res {
      msg = "записал"
      return msg
    }

    switch input {
      case "/help":
        msg = "описание о боте"
      case "/settings":
        msg = "описание настроек"
      default:
        msg = "команда непонятна"
    }

    return msg
}
