package bot_api

import (
  "fmt"

  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
  "app-telegram/models"
)


// func TestCreateResponse(t *testing.T) {
//   test.Setup()

//   Convey("should return string poloniex", t, func() {
//     res := CreateResponse("p SBD > 0.000020", 123)

//     So(res, ShouldEqual, "ok")
//   })

//   Convey("should return poloniex list", t, func() {
//     res := CreateResponse("plist", 123)

//     So(res, ShouldEqual, "poliniex list \n")
//   })
// }

func TestCreateAlert(t *testing.T) {
  test.Setup()
  var id_telegram int64

  // user not found
  Convey("should return string poloniex", t, func() {
    id_telegram = 123

    res := CreateAlert("p SBD > 0.000020", id_telegram)
    fmt.Println(11111)


    user, _ := models.FindByIdTelegramm(id_telegram)


    fmt.Println(999, user.Alerts)


    So(res, ShouldEqual, "ok")
  })

  // user exist
  // Convey("should return poloniex list", t, func() {
  //   res := CreateAlert("p SBD > 0.000020", 123)

  //   So(res, ShouldEqual, "list")
  // })
}
