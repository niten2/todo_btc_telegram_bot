package models

import (
  // "fmt"
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
  "app-telegram/db"
  . "app-telegram/request"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestCoin(t *testing.T) {
  Db := db.Connect()
  Db.DropDatabase()

  Convey("CreateCoin", t, func() {
    coin := CreateCoin(Db, "BTC_USDT", "0.00003")

    So(coin, ShouldNotBeNil)
  })

  Convey("FindCoin", t, func() {
    CreateCoin(Db, "BTC_USDT", "0.00003")

    coin := FindCoin(Db, "BTC_USDT")

    So(coin, ShouldNotBeNil)
  })

  Convey("BuildCoins", t, func() {
    coins := make(map[string]PoloniexCoin)
    coins["BTC_BCN"] = PoloniexCoin{Last: "0.00000066"}
    coins["BTC_BELA"] = PoloniexCoin{Last: "0.00003236"}

    res := BuildCoins(Db, coins)

    var values []Coin
    values = append(values, Coin{Name: "BTC_BCN", Value: "0.00000066"})
    values = append(values, Coin{Name: "BTC_BELA", Value: "0.00003236"})

    So(res, ShouldResemble, values)
  })

  Convey("CreateCoins", t, func() {
    var values []Coin
    values = append(values, Coin{Name: "BTC_BCN", Value: "0.00000066"})
    values = append(values, Coin{Name: "BTC_BELA", Value: "0.00003236"})

    coins := CreateCoins(Db, values)

    So(coins, ShouldNotBeNil)
  })

}
