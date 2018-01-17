package models

import (
  // "fmt"
  "testing"
  // "os"
  // "strconv"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
  "app-telegram/db"
)

func init() {
  _ = godotenv.Load("../.env.test")
}

func TestUser(t *testing.T) {
  Db := db.Connect()
  Db.DropDatabase()

  Convey("CreateUser", t, func() {
    _ = CreateUser(Db, "nam", "id-ttt")

    find_user := FindUser(Db, "nam")

    So(find_user, ShouldNotBeNil)
  })

}
