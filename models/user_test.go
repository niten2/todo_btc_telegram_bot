package models

import (
  // "fmt"
  // "os"
  // "strconv"
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
  "app-telegram/db"
)

func init() {
  _ = godotenv.Load("../.env.test")
	db.Connect()
}

func TestUser(t *testing.T) {
  db.Db.DropDatabase()

  Convey("CreateUser", t, func() {
    _ = CreateUser("nam", "id-ttt")

    find_user := FindUser("nam")

    So(find_user, ShouldNotBeNil)
  })

}
