package schedule

import (
  // "fmt"
  "gopkg.in/h2non/gock.v1"
	"gopkg.in/telegram-bot-api.v4"
  "log"
  "testing"
)

func TestSimple(t *testing.T) {
  defer gock.Off()

  gock.New("https://api.telegram.org/bottoken").
    Post("/getMe").
    Reply(200).
    JSON(map[string]string{"user": "test"})

  bot, err := tgbotapi.NewBotAPI("token")

  // !!!!
	if err != nil {
		log.Panic(err)
	}

  Run(bot)
}
