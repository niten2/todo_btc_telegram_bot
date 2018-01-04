package request

import (
  "fmt"
	"testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
)

func init() {
  _ = godotenv.Load("../.env")
}

func TestHelloWorld(t *testing.T) {

  Convey("should return balance", t, func() {
    res := GetBitapsBalanceWallet()

    So(res, ShouldBeGreaterThan, 200)
  })

}
