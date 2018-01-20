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
    res := Settings()

    ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))
    TelegramUserId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)

    So(res.DbUrl, ShouldEqual, os.Getenv("DB_URL"))
    So(res.DbName, ShouldEqual, "app_telegram_test")

    So(res.TelegramToken, ShouldEqual, os.Getenv("TELEGRAM_TOKEN"))
    So(res.TelegramUserId, ShouldEqual, TelegramUserId)

    So(res.WalletId, ShouldEqual, os.Getenv("WALLET_ID"))
    So(res.ScheduleEverySeconds, ShouldEqual, ScheduleEverySeconds)
  })

}
