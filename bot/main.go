package bot

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
  _ = godotenv.Load()

  telegram_token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(telegram_token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

  return bot
}

func SendMessage(bot *tgbotapi.BotAPI) {
  user_id, _ := strconv.ParseInt(os.Getenv("USER_ID"), 10, 64)
  message := CreateMessage()

  bot.Send(tgbotapi.NewMessage(user_id, message))

  fmt.Println("bot send message user_id %s", user_id)
}

func CreateMessage() string {
  usd := request.GetCoinmarketcapCurrentBtc()

  return fmt.Sprintf(`
    текущий курс %s \n
    другое значение ---
  `, usd)
}
