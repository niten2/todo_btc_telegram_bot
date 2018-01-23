package models

import (
  // "fmt"

  // "os"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
  // "app-telegram/db"
  // "app-telegram/types"
)

func TestUser(t *testing.T) {
  test.Setup()

  Convey("NewUser", t, func() {
    user := NewUser("test", 123)

    So(user, ShouldNotBeNil)
  })

  Convey("Create", t, func() {
    user := NewUser("test", 123)
    user.Create()

    user, _ = FindUserById(string(user.ID))

    So(user, ShouldNotBeNil)
  })

  Convey("Save", t, func() {
    user := CreateUser("test", 123)
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

  Convey("FindUserByIdTelegramm", t, func() {
    user := NewUser("test", 123)
    user.Create()
    user, _ = FindUserByIdTelegramm(123)
    So(user, ShouldNotBeNil)
  })

}
