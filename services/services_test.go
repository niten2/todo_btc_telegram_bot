package services

import (
  // "fmt"
	"testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestGenerateId(t *testing.T) {

  Convey("should", t, func() {
    res := GenerateId()

    So(res, ShouldNotBeNil)
  })

}
