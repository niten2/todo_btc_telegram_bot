package config

import (
  // "fmt"

	"testing"
  "os"
  "strconv"
  . "github.com/smartystreets/goconvey/convey"

)

func TestSettings(t *testing.T) {
  os.Setenv("ENV", "test")

  Convey("should return balance", t, func() {
    res := Settings()

    ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))
    TelegramUserId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)

    So(res.Env, ShouldEqual, "test")
    So(res.IsEnvTest, ShouldEqual, true)

    So(res.DbUrl, ShouldEqual, os.Getenv("DB_URL"))
    So(res.DbName, ShouldEqual, "app_telegram_test")

    So(res.TelegramToken, ShouldEqual, os.Getenv("TELEGRAM_TOKEN"))
    So(res.TelegramUserId, ShouldEqual, TelegramUserId)

    So(res.WalletId, ShouldEqual, os.Getenv("WALLET_ID"))
    So(res.ScheduleEverySeconds, ShouldEqual, ScheduleEverySeconds)
  })

}
