package models

import (
  "fmt"

  // "os"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
  // "app-telegram/db"
  // "app-telegram/types"
)

// func TestUser(t *testing.T) {
// 	db.Connect()
//   db.Db.DropDatabase()
//   os.Setenv("ENV", "test")

//   Convey("create user", t, func() {
//     user := User{Name: "test"}
//     user.Create()

//     find_user := FindUserByName("test")

//     So(find_user, ShouldNotBeNil)
//   })

//   Convey("update user", t, func() {
//     user := NewUser()

//     user.Name = "test"
//     user.Create()

//     user.Name = "new"
//     user.Update()

//     find_user := FindUserById(user.ID)

//     So(find_user.Name, ShouldEqual, "new")
//   })

// }

func TestCreatePoloniexAlert(t *testing.T) {
  test.Setup()
  test.SetDebug()

  // Convey("NewUser", t, func() {
  //   user := NewUser("test", 123)

  //   So(user, ShouldNotBeNil)
  // })

  // Convey("Create", t, func() {
  //   user := NewUser("test", 123)
  //   user.Create()

  //   user, _ = FindUserById(string(user.ID))

  //   So(user, ShouldNotBeNil)
  // })

  Convey("Save", t, func() {
    user := CreateUser("test", 123)
    alert, _ := NewAlert("p SBD > 0.000020")

    user.Alerts = append(user.Alerts, alert)
    user.Save()

    user, _ = FindUserById(string(user.ID))


    fmt.Println(user)

    fmt.Println(user.Alerts)

    // fmt.Println(111, user)

    // So(user, ShouldNotBeNil)
  })

  // Convey("FindUserById", t, func() {
  //   user := NewUser("test", 123)
  //   user.Create()
  //   user, _ = FindUserById(string(user.ID))
  //   So(user, ShouldNotBeNil)
  // })

  // Convey("FindUserByName", t, func() {
  //   user := NewUser("test", 123)
  //   user.Create()
  //   user, _ = FindUserByName("test")
  //   So(user, ShouldNotBeNil)
  // })

  // Convey("FindByIdTelegramm", t, func() {
  //   user := NewUser("test", 123)
  //   user.Create()
  //   user, _ = FindByIdTelegramm(123)
  //   So(user, ShouldNotBeNil)
  // })

}
