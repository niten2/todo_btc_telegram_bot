package schedule

import (
	// "fmt"
  "log"
	"github.com/jasonlvhit/gocron"
	"gopkg.in/telegram-bot-api.v4"
  "app-telegram/bot_api"
)

func Run (bot *tgbotapi.BotAPI) {
	// gocron.Every(2).Seconds().Do(bot_api.SendMessage, bot)
	gocron.Every(1).Day().At("10:30").Do(bot_api.SendMessage, bot)

  gocron.Start()

  log.Printf("start schedule every one seconds")
}

