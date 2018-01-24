package bot_api

import (
  // "fmt"

  "testing"
  . "github.com/smartystreets/goconvey/convey"
  "gopkg.in/h2non/gock.v1"

  "app-telegram/test"
)

func TestBotApi(t *testing.T) {
  test.Setup()

  Convey("should run InitBot", t, func() {
    gock.New("https://api.telegram.org/(.*)").
      Post("/getMe").
      Reply(200).
      BodyString(test.TelegramBotBodyString)
    defer gock.Off()

    InitBot()

    So(Bot, ShouldNotBeNil)
  })

}
