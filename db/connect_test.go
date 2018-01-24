package db

import (
  // "fmt"

  "os"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)


func TestConnect(t *testing.T) {
  os.Setenv("ENV", "test")

  Convey("should", t, func() {
    Connect()

    So(Db, ShouldNotBeNil)
    So(Session, ShouldNotBeNil)
  })
}
