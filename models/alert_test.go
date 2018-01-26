package models

import (
	// "fmt"

	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"app-telegram/test"
)

func TestAlert(t *testing.T) {
	test.Setup()

	Convey("NewAlert without coin", t, func() {
		res, err := NewAlert("p SBD > 0.000020")

		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
	})

	Convey("NewAlert with coin", t, func() {
		CreateCoin("BTC_SBD", 1, 1)
		res, err := NewAlert("p sbd > 0.000020")

		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		So(res.Name, ShouldEqual, "BTC_SBD")
	})
}
