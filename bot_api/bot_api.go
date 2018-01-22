package bot_api

import (
  "fmt"
  "gopkg.in/telegram-bot-api.v4"

  "app-telegram/request"
  "app-telegram/config"

  . "app-telegram/logger"
  "github.com/sirupsen/logrus"
)

func InitBot() *tgbotapi.BotAPI {
  telegram_token := config.Settings().TelegramToken
  bot, err := tgbotapi.NewBotAPI(telegram_token)

  if err != nil {
    Log.Fatal(err)
  }

  // bot.Debug = true

  Log.WithFields(logrus.Fields{
    "user": bot.Self.UserName,
  }).Info("Authorized on account")

  return bot
}

func CreateMessage() string {
  usd := request.GetCoinmarketcapCurrentBtc()
  balance_wallet := request.GetBitapsBalanceWallet()

  return fmt.Sprintf(`
    текущий курс %s
    btc в кошельке 0.%d
  `, usd, balance_wallet)
}

func SendMessage(bot *tgbotapi.BotAPI) {
  user_id := config.Settings().TelegramUserId

  message := CreateMessage()

  bot.Send(tgbotapi.NewMessage(user_id, message))

  Log.WithFields(logrus.Fields{
    "user_id": user_id,
    "message": message,
  }).Info("send message")
}
