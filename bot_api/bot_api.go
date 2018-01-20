package bot_api

import (
  "fmt"
  // "os"
  "log"
  // "strconv"
  "gopkg.in/telegram-bot-api.v4"

  "app-telegram/request"
  "app-telegram/config"
)

func InitBot() *tgbotapi.BotAPI {
  telegram_token := config.Settings().TelegramToken
  bot, err := tgbotapi.NewBotAPI(telegram_token)

  if err != nil {
    log.Panic(err)
  }

  // bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

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

  fmt.Println("bot send message user_id %s", user_id)
}
