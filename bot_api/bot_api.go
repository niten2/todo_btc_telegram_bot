package bot_api

import (
	// "fmt"

	"app-telegram/logger"
	"github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"

	"app-telegram/config"
)

var Bot *tgbotapi.BotAPI

func InitBot() {
	telegram_token := config.Settings().TelegramToken

	if telegram_token == "" {
		logger.Log.Fatal("telegram_token not found")
	}

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
