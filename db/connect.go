package db

import (
  "fmt"
  // "os"
  "gopkg.in/mgo.v2"
  "app-telegram/config"

  . "app-telegram/logger"
  "github.com/sirupsen/logrus"

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

  if db_url == "" {
    Log.Fatal("db_url not found")
  }

  mongo, err := mgo.ParseURL(db_url)

  s, err := mgo.Dial(db_url)

  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    panic(err.Error())
  }

  s.SetSafe(&mgo.Safe{})

  Log.WithFields(logrus.Fields{
    "db_url": db_url,
  }).Info("connect mongo")

  Session = s
  Mongo = mongo
  Db = s.DB(db_name)
}
