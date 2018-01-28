package config

import (
	// "fmt"

	. "github.com/smartystreets/goconvey/convey"
	"os"
	// "strconv"
	"testing"
)

func TestSettings(t *testing.T) {
	os.Setenv("ENV", "test")

	Convey("should return balance", t, func() {
		res := Settings()

		So(res.Env, ShouldEqual, "test")
		So(res.IsEnvTest, ShouldEqual, true)

		So(res.DbUrl, ShouldEqual, os.Getenv("DB_URL"))
		So(res.DbName, ShouldEqual, "app_telegram_test")

		So(res.TelegramToken, ShouldEqual, os.Getenv("TELEGRAM_TOKEN"))
	})

}
