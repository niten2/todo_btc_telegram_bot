package bot_api

import (
  // "fmt"

  "gopkg.in/telegram-bot-api.v4"
  "app-telegram/logger"
  "github.com/sirupsen/logrus"

  "app-telegram/config"
)

var Bot *tgbotapi.BotAPI

func InitBot() {
  telegram_token := config.Settings().TelegramToken
  bot, err := tgbotapi.NewBotAPI(telegram_token)

  if err != nil {
    logger.Log.Fatal(err)
  }

  // bot.Debug = true

  logger.Log.WithFields(logrus.Fields{
    "user": bot.Self.UserName,
  }).Info("Authorized on account")

  Bot = bot
}

