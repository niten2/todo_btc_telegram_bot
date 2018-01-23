package bot_api

import (
  "fmt"

  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
  "app-telegram/models"
)

func TestCreateResponse(t *testing.T) {
  test.Setup()

  Convey("should return string poloniex", t, func() {
    res := CreateResponse("p SBD > 0.000020", 123)

    So(res, ShouldEqual, "ok")
  })

  Convey("should return poloniex list", t, func() {
    res := CreateResponse("plist", 123)

    So(res, ShouldEqual, "poliniex list \n")
  })
}

func TestCreateAlert(t *testing.T) {
  test.Setup()
  var id_telegram int64

  Convey("user not found, should return string poloniex", t, func() {
    id_telegram = 123

    res := CreateAlert("p SBD > 0.000020", id_telegram)
    user, _ := models.FindUserByIdTelegram(id_telegram)

    fmt.Println(999, user.Alerts)

    So(res, ShouldEqual, "ok")

    alert := user.Alerts[0]

    So(alert.Name, ShouldEqual, "BTC_SBD")
    So(alert.Compare, ShouldEqual, ">")
    So(alert.Value, ShouldEqual, 0.000020)
  })
}

func TestCreatePoloniexCoinList(t *testing.T) {
  test.Setup()

  Convey("should return list poloniex", t, func() {
    _, _ = models.CreateCoin("test", 123)

    res := CreatePoloniexCoinList()

    So(res, ShouldContainSubstring, "test")
    So(res, ShouldContainSubstring, "123")
  })

  Convey("should return list poloniex", t, func() {
    test.DropDatabase()

    res := CreatePoloniexCoinList()

    fmt.Println(res)

    So(res, ShouldNotContainSubstring, "list")
  })
}
