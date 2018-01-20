package main

import (
  // "fmt"
  "os"
  "testing"

  "gopkg.in/h2non/gock.v1"
  . "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
  os.Setenv("ENV", "test")

  Convey("should", t, func() {

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


    // bot := bot_api.InitBot()

    // fmt.Println(bot)


    // fmt.Println("before")

    // gock.New("(.*)").
    //   Post("(.*)").
    //   Reply(200).
    //   JSON(map[string]string{"user": "test"})

    main()

    // Connect()

    // So(Db, ShouldNotBeNil)
    // So(Session, ShouldNotBeNil)
  })

}
