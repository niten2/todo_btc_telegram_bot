package schedule

// import (
//   "fmt"
//   "gopkg.in/h2non/gock.v1"
// 	"gopkg.in/telegram-bot-api.v4"
//   // "log"
//   "testing"
// )

// func TestSchedule(t *testing.T) {
//   defer gock.Off()

//   defer gock.Disable()

//   gock.New("(.*)").
//     Post("(.*)").
//     Reply(200).
//     JSON(map[string]string{"user": "test"})

//   bot, err := tgbotapi.NewBotAPI("token")

//   fmt.Println(err)
//   fmt.Println(bot)

//   // !!!!
// 	if err != nil {
//     fmt.Println(1111)
//     fmt.Println(err)


// 		log.Panic(err.Error())
// 	}

//   Run(bot)
// }
