package db

import (
  // "fmt"
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestConnect(t *testing.T) {

  Convey("should", t, func() {
    db, session := Connect()
    defer session.Close()

    So(db, ShouldNotBeNil)
    So(session, ShouldNotBeNil)
  })

}
