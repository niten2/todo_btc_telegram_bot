package logger

import (
  // "fmt"

  "os"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

)

func TestInitFileLogger(t *testing.T) {
  os.Setenv("ENV", "test")

  Convey("should log.out exist", t, func() {
    InitFileLogger()

    So(Log.Out, ShouldNotBeNil)
  })

}
