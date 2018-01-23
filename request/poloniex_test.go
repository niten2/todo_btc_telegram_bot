package request

import (
  // "fmt"

	"gopkg.in/h2non/gock.v1"
	"testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
)

func TestRequestPoloniex(t *testing.T) {
  test.Setup()

  Convey("should return valid result", t, func() {

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

    res, _ := RequestPoloniex()

    So(res, ShouldContain, CoinPoloniex{
      Name: "BTC_BCN",
      Value: 0.00000066,
    })

    So(res, ShouldContain, CoinPoloniex{
      Name: "BTC_BELA",
      Value: 0.00003236,
    })
  })
}
