package models

import (
  // "fmt"

  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "app-telegram/test"
)


func TestAlert(t *testing.T) {
  test.Setup()

  Convey("NewAlert", t, func() {
    res, _ := NewAlert("p SBD > 0.000020")

    So(res, ShouldNotBeNil)
  })
}
