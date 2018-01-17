package bot_api

import (
  // "fmt"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestCreateMessage(t *testing.T) {

  Convey("should ", t, func() {
    res := CreateResponse("p BTC SBD 0.000020")

    So(res, ShouldEqual, "записал")
  })

}
