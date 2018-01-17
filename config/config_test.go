package config

import (
  // "fmt"
	"testing"
  "os"
  "github.com/joho/godotenv"
  "strconv"
  . "github.com/smartystreets/goconvey/convey"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestSettings(t *testing.T) {

  Convey("should return balance", t, func() {
    ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

    res := Settings()

    So(res.DbUrl, ShouldEqual, os.Getenv("DB_URL"))
    So(res.DbName, ShouldEqual, "app_telegram_test")

    So(res.TelegramToken, ShouldEqual, os.Getenv("TELEGRAM_TOKEN"))
    So(res.TelegramUserId, ShouldEqual, os.Getenv("TELEGRAM_USER_ID"))
    So(res.WalletId, ShouldEqual, os.Getenv("WALLET_ID"))
    So(res.ScheduleEverySeconds, ShouldEqual, ScheduleEverySeconds)
  })

}
