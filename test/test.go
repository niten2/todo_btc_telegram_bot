package test

import (
  // "fmt"

  "os"
  "log"
  "gopkg.in/mgo.v2"
  // "gopkg.in/h2non/gock.v1"

  "app-telegram/db"
)

const TelegramBotBodyString = `
  {
    "ok": true,
    "result": {
      "id": 534094005,
      "is_bot": true,
      "first_name": "CoinInfo",
      "username": "coint_info_bot"
    }
  }
`

func Setup() {
  os.Setenv("ENV", "test")

	db.Connect()
  db.Db.DropDatabase()
}

func SetupEnv() {
  os.Setenv("ENV", "test")
}

func DropDatabase() {
  db.Db.DropDatabase()
}

func SetDebugDb() {
  mgo.SetDebug(true)
  var aLogger *log.Logger
  aLogger = log.New(os.Stderr, "", log.LstdFlags)
  mgo.SetLogger(aLogger)
}
