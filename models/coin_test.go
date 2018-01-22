package models

import (
  // "fmt"

  // "os"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  // . "app-telegram/request"
  // "app-telegram/db"
  "app-telegram/test"
)


func TestCoin(t *testing.T) {
  test.Setup()

  // Convey("CreateCoin", t, func() {
  //   coin := CreateCoin("BTC_USDT", "0.00003")

  //   So(coin, ShouldNotBeNil)
  // })

  // Convey("FindCoin", t, func() {
  //   CreateCoin("BTC_USDT", "0.00003")
  //   coin := FindCoin("BTC_USDT")

  //   So(coin, ShouldNotBeNil)
  // })

  // Convey("BuildCoins", t, func() {
  //   coins := make(map[string]PoloniexCoin)
  //   coins["BTC_BCN"] = PoloniexCoin{Last: "0.00000066"}
  //   coins["BTC_BELA"] = PoloniexCoin{Last: "0.00003236"}

  //   res := BuildCoins(coins)

  //   var values []Coin
  //   values = append(values, Coin{Name: "BTC_BCN", Value: "0.00000066"})
  //   values = append(values, Coin{Name: "BTC_BELA", Value: "0.00003236"})

  //   So(res, ShouldResemble, values)
  // })

  // Convey("UpdateCoinsPoloniex", t, func() {
  //   // var values []Coin
  //   // values = append(values, Coin{Name: "BTC_BCN", Value: "0.00000066"})
  //   // values = append(values, Coin{Name: "BTC_BELA", Value: "0.00003236"})

  //   UpdateCoinsPoloniex()

  //   // So(coins, ShouldNotBeNil)
  // })

  // Convey("CreatePoloniexCoinList", t, func() {
  //   CreateCoin("BTC_USDT", "0.00003")
  //   CreateCoin("BTC_XRP", "0.00001")

  //   res := CreatePoloniexCoinList()

  //   So(res, ShouldNotBeNil)
  // })

  Convey("CreatePoloniexAlert", t, func() {
    // CreateCoin("BTC_SBD", 0.00003)

    // alert := NewAlert("p SBD > 0.000020")

    // res := CreatePoloniexAlert(alert)

    // So(res, ShouldNotBeNil)
  })
}
