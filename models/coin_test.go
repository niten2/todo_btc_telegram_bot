package models

import (
  // "fmt"

  "testing"
  . "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"

  "app-telegram/test"
)


func TestCoin(t *testing.T) {
  test.Setup()

  Convey("Coin.Create()", t, func() {
    coin := NewCoin("test", 123)
    coin.Create()

    coin, _ = FindCoinById(string(coin.ID))

    So(coin, ShouldNotBeNil)
  })

  Convey("Coin.Save()", t, func() {
    coin, _ := CreateCoin("test", 123)
    coin, _ = FindCoinById(string(coin.ID))

    So(coin.Name, ShouldEqual, "test")
    So(coin.Value, ShouldEqual, 123)
  })

  Convey("NewCoin", t, func() {
    coin := NewCoin("test", 123)

    So(coin, ShouldNotBeNil)
    So(coin.ID, ShouldNotBeNil)
  })

  Convey("FindCoinById", t, func() {
    coin, _ := CreateCoin("test", 123)
    coin, _ = FindCoinById(string(coin.ID))

    So(coin, ShouldNotBeNil)
  })

  Convey("FindCoinByName", t, func() {
    coin, _ := CreateCoin("test", 123)
    coin, _ = FindCoinByName(string(coin.Name))

    So(coin, ShouldNotBeNil)
  })

  Convey("FetchCoin", t, func() {
    body := `{
      "BTC_BCN": {
        "id": 7,
        "last": "0.00000066",
        "lowestAsk": "0.00000067",
        "highestBid": "0.00000066",
        "percentChange": "-0.14285714",
        "baseVolume": "222.79844664",
        "quoteVolume": "317319368.28207934",
        "isFrozen": "0",
        "high24hr": "0.00000077",
        "low24hr": "0.00000065"
      },
      "BTC_BELA": {
        "id": 8,
        "last": "0.00003236",
        "lowestAsk": "0.00003239",
        "highestBid": "0.00003200",
        "percentChange": "-0.08094291",
        "baseVolume": "18.80739029",
        "quoteVolume": "570326.49484689",
        "isFrozen": "0",
        "high24hr": "0.00003586",
        "low24hr": "0.00003200"
      }
    }`

    defer gock.Disable()

    gock.New("https://poloniex.com/public?command=returnTicker").
      Reply(200).
      BodyString(body)

    _ = FetchCoin()

    coin1, _ := FindCoinByName("BTC_BCN")
    coin2, _ := FindCoinByName("BTC_BELA")

    So(coin1, ShouldNotBeNil)
    So(coin2, ShouldNotBeNil)
  })

}
