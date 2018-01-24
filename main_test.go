package main

import (
  // "fmt"

  // "os"
  "testing"

  "gopkg.in/h2non/gock.v1"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
)

func TestMain(t *testing.T) {
  test.Setup()

  Convey("should call main", t, func() {
    gock.New("https://api.telegram.org/(.*)").
      Post("/getMe").
      Reply(200).
      BodyString(test.TelegramBotBodyString)
    defer gock.Off()

    main()
  })

}
