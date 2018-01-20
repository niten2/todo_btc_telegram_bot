package db

import (
  "fmt"
  // "os"
  "gopkg.in/mgo.v2"
  "app-telegram/config"

  // log "github.com/sirupsen/logrus"
)

var (
  Session *mgo.Session
  Mongo *mgo.DialInfo
  Db *mgo.Database
)

func Connect() {
  db_url := config.Settings().DbUrl
  db_name := config.Settings().DbName

  mongo, err := mgo.ParseURL(db_url)
  // _, err := mgo.ParseURL(db_url)

  s, err := mgo.Dial(db_url)

  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    panic(err.Error())
  }

  s.SetSafe(&mgo.Safe{})

  // log.Info("A walrus appears")

  // fmt.Println("Connected to mongo", db_url)

  Session = s
  Mongo = mongo
  Db = s.DB(db_name)
}
