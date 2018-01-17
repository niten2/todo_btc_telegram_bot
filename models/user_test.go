package models

import (
  "fmt"
  "testing"
  // "os"
  // "strconv"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
  "app-telegram/db"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestConnect(t *testing.T) {

  Convey("should", t, func() {
    fmt.Println(111)

    // ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

    db.Connect()

    user := CreateUser("nameeee")

    fmt.Println(user)


    // fmt.Println(Db)

    // fmt.Println(DbUrl)


    // Connect()

    // res := Connect()

    // fmt.Println(res)


    // So(res.DbUrl, ShouldEqual, os.Getenv("DB_URL"))
    // So(res.TelegramToken, ShouldEqual, os.Getenv("TELEGRAM_TOKEN"))
    // So(res.TelegramUserId, ShouldEqual, os.Getenv("TELEGRAM_USER_ID"))
    // So(res.WalletId, ShouldEqual, os.Getenv("WALLET_ID"))
    // So(res.ScheduleEverySeconds, ShouldEqual, ScheduleEverySeconds)
  })

}
