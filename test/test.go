package test

import (
  // "fmt"
  // "testing"
  // "github.com/joho/godotenv"
  // . "github.com/smartystreets/goconvey/convey"
  "os"
  "log"
  "gopkg.in/mgo.v2"
  "app-telegram/db"
)

func SetupEnv() {
  os.Setenv("ENV", "test")
}

func Setup() {
  os.Setenv("ENV", "test")

	db.Connect()
  db.Db.DropDatabase()
}

func SetDebugDb() {
  mgo.SetDebug(true)
  var aLogger *log.Logger
  aLogger = log.New(os.Stderr, "", log.LstdFlags)
  mgo.SetLogger(aLogger)
}
