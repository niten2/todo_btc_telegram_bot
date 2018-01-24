package schedule

import (
	"fmt"

	"github.com/jasonlvhit/gocron"

  "app-telegram/bot_api"
)

func InitSchedule() {
  // NOTE examples
  // gocron.Every(2).Seconds().Do(bot_api.SendMessage, bot)
  // gocron.Every(1).Day().At("10:00").Do(bot_api.SendMessage, bot)
  // gocron.Every(10).Minutes().Do(models.UpdateCoinsPoloniex)

  fmt.Println(2222)

  gocron.Every(1).Minutes().Do(bot_api.CheckCoin)

  gocron.Start()
}
