package main

import (
  "fmt"
  "testing"
  "github.com/joho/godotenv"
  "gopkg.in/h2non/gock.v1"

  . "github.com/smartystreets/goconvey/convey"
)

func init() {
  _ = godotenv.Load("./.env.test")
}

func TestMain(t *testing.T) {

  Convey("should", t, func() {

    // bot := bot_api.InitBot()

    // fmt.Println(bot)


    // fmt.Println("before")

    // gock.New("(.*)").
    //   Post("(.*)").
    //   Reply(200).
    //   JSON(map[string]string{"user": "test"})

    // main()

    // Connect()

    // So(Db, ShouldNotBeNil)
    // So(Session, ShouldNotBeNil)
  })

}
