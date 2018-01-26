package bot_api

import (
	// "fmt"

	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"app-telegram/models"
	"app-telegram/test"
)

func TestCreateResponse(t *testing.T) {
	test.Setup()

	Convey("should return string poloniex", t, func() {
		models.CreateCoin("BTC_SBD", 1)
		res := CreateResponse("p SBD > 0.000020", 123)

		So(res, ShouldEqual, "Alert successfully added")
	})

	Convey("should return string poloniex", t, func() {
		test.DropDatabase()
		res := CreateResponse("p SBD > 0.000020", 123)

		So(res, ShouldEqual, "Something went wrong")
	})

	Convey("should return poloniex list", t, func() {
		test.DropDatabase()
		res := CreateResponse("plist", 123)

		So(res, ShouldEqual, "Data not found")
	})
}

func TestCreateAlert(t *testing.T) {
	test.Setup()
	var id_telegram int64

	Convey("user not found, should return string poloniex", t, func() {
		id_telegram = 123
		models.CreateCoin("BTC_SBD", 1)

		res := CreateAlert("p SBD > 0.000020", id_telegram)
		user, _ := models.FindUserByIdTelegram(id_telegram)

		So(res, ShouldEqual, "Alert successfully added")

		alert := user.Alerts[0]

		So(alert.Name, ShouldEqual, "BTC_SBD")
		So(alert.Compare, ShouldEqual, ">")
		So(alert.Value, ShouldEqual, 0.000020)
	})

	Convey("coin not found, user not found, should return string poloniex", t, func() {
		test.DropDatabase()
		id_telegram = 123

		res := CreateAlert("p SBD > 0.000020", id_telegram)

		So(res, ShouldEqual, "Something went wrong")
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

		So(res, ShouldContainSubstring, "Data not found")
	})
}

func TestCheckUsersAlert(t *testing.T) {
	test.Setup()

	Convey("should return list poloniex", t, func() {
		user, _ := models.CreateUser("user", 123)
		user.AddAlert("p SBD < 0.1")
		models.CreateCoin("BTC_SBD", 0.02)

		_, _ = models.CreateCoin("BTC_XRP", 1.5)

		CheckUsersAlert()

		user, _ = models.FindUserById(string(user.ID))

		So(user.Alerts, ShouldBeEmpty)
	})
}

func TestAddUsdtCoin(t *testing.T) {
	test.Setup()

	Convey("should return list poloniex", t, func() {
		models.CreateCoin("BTC_SBD", 3)
		models.CreateCoin("USDT_BTC", 2.5)

		AddUsdtCoin()

		coin, _ := models.FindCoinByName("BTC_SBD")

		So(coin.ValueUsdt, ShouldEqual, 7.5)
	})
}
