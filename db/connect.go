package db

import (
  "fmt"
  // "os"
  "gopkg.in/mgo.v2"
  "app-telegram/config"
)

var (
  Session *mgo.Session
  Mongo *mgo.DialInfo
  Db *mgo.Database
)

func Connect() {
  db_url := config.Settings().DbUrl

  mongo, err := mgo.ParseURL(db_url)

  s, err := mgo.Dial(db_url)

  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    panic(err.Error())
  }

  s.SetSafe(&mgo.Safe{})

  fmt.Println("Connected to mongo", db_url)

  Session = s
  Mongo = mongo
  Db = s.DB("test")
}
