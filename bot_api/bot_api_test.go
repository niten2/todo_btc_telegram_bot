package bot_api

import (
  // "fmt"
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"

  "gopkg.in/h2non/gock.v1"
)

func init() {
  _ = godotenv.Load("../.env")
}

func TestBotApi(t *testing.T) {
  Convey("should run InitBot", t, func() {
    defer gock.Off()

    gock.New("https://api.telegram.org/(.*)").
      Post("/getMe").
      Reply(200).
      BodyString(`
        {
          "ok": true,
          "result": {
            "id": 534094005,
            "is_bot": true,
            "first_name": "CoinInfo",
            "username": "coint_info_bot"
          }
        }
      `)

    res := InitBot()

    So(res, ShouldNotBeNil)
  })

  // TODO need stub requests
  // Convey("should run CreateMessage", t, func() {
  //   res := CreateMessage()

  //   So(res, ShouldNotBeBlank)
  // })

}
