package models

import (
	"fmt"

	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"app-telegram/test"
)

func TestUser(t *testing.T) {
	test.Setup()

	Convey("Create", t, func() {
		user := NewUser("test", 123)
		user.Create()

		user, _ = FindUserById(string(user.ID))

		So(user, ShouldNotBeNil)
	})

	Convey("Save", t, func() {
		user, _ := CreateUser("test", 123)
		alert, _ := NewAlert("p SBD > 0.000020")

		user.Alerts = append(user.Alerts, alert)
		user.Save()

		user, _ = FindUserById(string(user.ID))

		name := user.Alerts[0].Name
		compare := user.Alerts[0].Compare
		value := user.Alerts[0].Value

		So(name, ShouldEqual, "BTC_SBD")
		So(compare, ShouldEqual, ">")
		So(value, ShouldEqual, 2e-05)
	})

	Convey("NewUser", t, func() {
		user := NewUser("test", 123)

		So(user, ShouldNotBeNil)
	})

	Convey("FindUserById", t, func() {
		user := NewUser("test", 123)
		user.Create()
		user, _ = FindUserById(string(user.ID))
		So(user, ShouldNotBeNil)
	})

	Convey("FindUserByName", t, func() {
		user := NewUser("test", 123)
		user.Create()
		user, _ = FindUserByName("test")
		So(user, ShouldNotBeNil)
	})

	Convey("FindUserByIdTelegram", t, func() {
		user := NewUser("test", 123)
		user.Create()
		user, _ = FindUserByIdTelegram(123)
		So(user, ShouldNotBeNil)
	})

	Convey("FindOrCreateUserByIdTelegram user not exist", t, func() {
		user, _ := FindOrCreateUserByIdTelegram("test", 123)

		So(user.Name, ShouldEqual, "test")
		So(user.IdTelegram, ShouldEqual, 123)
	})

	Convey("FindOrCreateUserByIdTelegram user exist", t, func() {
		_, _ = CreateUser("test", 123)

		user, _ := FindOrCreateUserByIdTelegram("new", 123)

		So(user.Name, ShouldEqual, "test")
		So(user.IdTelegram, ShouldEqual, 123)
	})
}

func TestCheckAndRemoveUserAlert(t *testing.T) {
	test.Setup()

	Convey("should run valid alert >", t, func() {
		user, _ := CreateUser("test", 123)

		user.AddAlert("p SBD > 100")

		CreateCoin("BTC_SBD", 100, 100)

		fmt.Println(111)

		message, _ := user.CheckAndRemoveUserAlert()

		resMessage := `ALERT BTC_SBD 100.000000 > 100.000000 \n`

		So(message, ShouldEqual, resMessage)
	})

	Convey("should run valid alert <", t, func() {
		user, _ := CreateUser("test", 123)
		user.AddAlert("p xpm < 1")

		CreateCoin("BTC_XPM", 0.00009103, 0.4)

		message, _ := user.CheckAndRemoveUserAlert()

		resMessage := `ALERT BTC_XPM 1.000000 < 0.400000 \n`

		So(message, ShouldEqual, resMessage)
	})

}
