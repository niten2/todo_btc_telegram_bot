package schedule

import (
	// "fmt"

	"github.com/jasonlvhit/gocron"

	"app-telegram/bot_api"
)

func InitSchedule() {
	gocron.Every(5).Minutes().Do(bot_api.CheckCoin)

	gocron.Start()
}
