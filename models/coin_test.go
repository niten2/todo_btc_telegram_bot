package models

import (
  // "fmt"
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
  "app-telegram/db"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestCoin(t *testing.T) {
  Db := db.Connect()
  Db.DropDatabase()

  Convey("CreateCoin", t, func() {
    CreateCoin(Db, "BTC_USDT", "0.00003")

    coin := FindCoin(Db, "BTC_USDT")

    So(coin, ShouldNotBeNil)
  })

}
