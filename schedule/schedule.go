package schedule

import (
	// "fmt"
  // "log"
  // "os"
  // "strconv"
  // "app-telegram/bot_api"
  // "app-telegram/models"

	"gopkg.in/telegram-bot-api.v4"
	"github.com/jasonlvhit/gocron"
)

func Run (bot *tgbotapi.BotAPI) {
  // schedule_every_seconds, err := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

	// if err != nil {
	// 	log.Panic(err)
	// }

  // if schedule_every_seconds {
  //   gocron.Every(2).Seconds().Do(bot_api.SendMessage, bot)
  //   log.Printf(`start schedule every 2 seconds`)
  // } else {
  //   gocron.Every(1).Day().At("10:00").Do(bot_api.SendMessage, bot)
  //   log.Printf(`start schedule every day at 10:00`)
  // }

  // log.Printf(`start schedule update poloniex coins every 10 minutes`)
  // gocron.Every(10).Minutes().Do(models.UpdateCoinsPoloniex)

  gocron.Start()
}

