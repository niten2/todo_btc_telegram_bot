package bot_api

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
	"log"
	"gopkg.in/telegram-bot-api.v4"
  "strconv"
  "app-telegram/request"
)

func InitBot() *tgbotapi.BotAPI {
  telegram_token := os.Getenv("TELEGRAM_TOKEN")
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
  user_id, _ := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)
  message := CreateMessage()

  bot.Send(tgbotapi.NewMessage(user_id, message))

  fmt.Println("bot send message user_id %s", user_id)
}
