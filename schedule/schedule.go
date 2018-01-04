package schedule

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"gopkg.in/telegram-bot-api.v4"
  "app-telegram/bot"
)

func Run (instanseBot *tgbotapi.BotAPI) {
	gocron.Every(5).Seconds().Do(bot.SendMessage, instanseBot)
	// gocron.Every(1).Day().At("10:30").Do(task)

	_, time := gocron.NextRun()
	fmt.Println(time)

  gocron.Start()
}

