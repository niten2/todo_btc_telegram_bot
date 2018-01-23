package bot_api

import (
  "fmt"

  "os"
  "regexp"
  "gopkg.in/telegram-bot-api.v4"

  "app-telegram/logger"
  "github.com/sirupsen/logrus"

  "app-telegram/models"
  "app-telegram/config"
)

const MessageError = "Что то пошло не так"

func InitActions () {
  fmt.Println("InitActions for bot")

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := Bot.GetUpdatesChan(u)

  if err != nil {
    logger.Log.Fatal(err)
  }

  // NOTE for testing
  if config.Settings().Env == "test" {
    os.Exit(0)
  }

  for update := range updates {
    if update.Message == nil {
      continue
    }

    user_name := update.Message.From.UserName
    id_telegram := update.Message.Chat.ID
    text := update.Message.Text

    _, _ = models.FindOrCreateUserByIdTelegram(user_name, id_telegram)

    logger.Log.WithFields(logrus.Fields{
      "user_name": user_name,
      "text": text,
      "id_telegram": id_telegram,
    }).Info("bot receive message")

    message := CreateResponse(text, id_telegram)
    response := tgbotapi.NewMessage(id_telegram, message)

    Bot.Send(response)
  }
}

func CreateResponse(input string, id_telegram int64) string {
  var msg string

  switch {
    case regexp.MustCompile(`^[p] [\D]* [\d\.]*`).MatchString(input):
      msg = CreateAlert(input, id_telegram)
    case regexp.MustCompile("^plist").MatchString(input):
      msg = CreatePoloniexCoinList()
    case regexp.MustCompile("/help").MatchString(input):
      msg = "описание о боте"
    case regexp.MustCompile("/settings").MatchString(input):
      msg = "описание настроек"
    default:
      msg = "команда непонятна"
  }

  return msg
}

func CreateAlert(input string, id_telegram int64) string {
  user, err := models.FindUserByIdTelegram(id_telegram)

  if err != nil {
    user = models.NewUser("name", id_telegram)
    user.Create()
  }

  alert, err := models.NewAlert(input)

  if err != nil {
    logger.Log.Warn(err)
    return MessageError
  }

  user.Alerts = append(user.Alerts, alert)
  err = user.Save()

  if err != nil {
    logger.Log.Warn(err)
    return MessageError
  }

  return "ok"
}

func CreatePoloniexCoinList() string {
  coins, err := models.FindCoinAll()

  if err != nil {
    logger.Log.Warn(err)
    return MessageError
  }

  if len(coins) == 0 {
    return "данные в базе отсуствуют"
  }

  return models.CreatePoloniexCoinList(coins)
}

func SendMessage(id_telegram int64, message string) {
  response := tgbotapi.NewMessage(id_telegram, message)
  Bot.Send(response)
}

func SendResponseError(id_telegram int64) {
  response := tgbotapi.NewMessage(id_telegram, MessageError)
  Bot.Send(response)
}

// NOTE check coin
func CheckCoin() {
  models.FetchCoin()
  CheckUsersAlert()
}

func CheckUsersAlert() {
  users, err := models.FindUserAll()

  if err != nil {
    logger.Log.Warn(err)
  }

  for _, user := range users {
    message, err := user.CheckAndRemoveUserAlert()

    if err != nil {
      logger.Log.Warn(err)
    }

    if message != "" {
      SendMessage(user.IdTelegram, message)
    }
  }
}
