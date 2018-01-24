package models

// import (
//   // "fmt"

//   "testing"
//   . "github.com/smartystreets/goconvey/convey"

//   "app-telegram/test"
// )

// func TestUser(t *testing.T) {
//   test.Setup()

//   Convey("Create", t, func() {
//     user := NewUser("test", 123)
//     user.Create()

//     user, _ = FindUserById(string(user.ID))

//     So(user, ShouldNotBeNil)
//   })

//   Convey("Save", t, func() {
//     user, _ := CreateUser("test", 123)
//     alert, _ := NewAlert("p SBD > 0.000020")

//     user.Alerts = append(user.Alerts, alert)
//     user.Save()

//     user, _ = FindUserById(string(user.ID))

//     name := user.Alerts[0].Name
//     compare := user.Alerts[0].Compare
//     value := user.Alerts[0].Value

//     So(name, ShouldEqual, "BTC_SBD")
//     So(compare, ShouldEqual, ">")
//     So(value, ShouldEqual, 2e-05)
//   })

//   Convey("NewUser", t, func() {
//     user := NewUser("test", 123)

//     So(user, ShouldNotBeNil)
//   })

//   Convey("FindUserById", t, func() {
//     user := NewUser("test", 123)
//     user.Create()
//     user, _ = FindUserById(string(user.ID))
//     So(user, ShouldNotBeNil)
//   })

//   Convey("FindUserByName", t, func() {
//     user := NewUser("test", 123)
//     user.Create()
//     user, _ = FindUserByName("test")
//     So(user, ShouldNotBeNil)
//   })

//   Convey("FindUserByIdTelegram", t, func() {
//     user := NewUser("test", 123)
//     user.Create()
//     user, _ = FindUserByIdTelegram(123)
//     So(user, ShouldNotBeNil)
//   })

//   Convey("FindOrCreateUserByIdTelegram user not exist", t, func() {
//     user, _ := FindOrCreateUserByIdTelegram("test", 123)

//     So(user.Name, ShouldEqual, "test")
//     So(user.IdTelegram, ShouldEqual, 123)
//   })

//   Convey("FindOrCreateUserByIdTelegram user exist", t, func() {
//     _, _ = CreateUser("test", 123)

//     user, _ := FindOrCreateUserByIdTelegram("new", 123)

//     So(user.Name, ShouldEqual, "test")
//     So(user.IdTelegram, ShouldEqual, 123)
//   })

//   Convey("CheckAndRemoveUserAlert", t, func() {
//     user, _ := CreateUser("test", 123)
//     user.AddAlert("p SBD > 0.000020")
//     CreateCoin("BTC_SBD", 100)

//     message, _ := user.CheckAndRemoveUserAlert()

//     resMessage := "ALERT BTC_SBD 0.000020 > 100.000000 \n"

//     So(message, ShouldEqual, resMessage)
//   })

//   Convey("RemoveAlert", t, func() {
//     user, _ := CreateUser("test", 123)
//     user.AddAlert("p SBD > 0.000020")
//     user.AddAlert("p SBC < 0.000020")

//     user.RemoveAlert(user.Alerts[0])

//     So(len(user.Alerts), ShouldEqual, 1)
//   })

// }
