package db

import (
  "fmt"
  // "os"
  "gopkg.in/mgo.v2"
  "app-telegram/config"
)

// var (
//   Session *mgo.Session
//   Mongo *mgo.DialInfo
//   Db *mgo.Database
// )

func Connect() (*mgo.Database, *mgo.Session) {
  db_url := config.Settings().DbUrl
  db_name := config.Settings().DbName

  // mongo, err := mgo.ParseURL(db_url)
  _, err := mgo.ParseURL(db_url)

  Session, err := mgo.Dial(db_url)

  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    panic(err.Error())
  }

  Session.SetSafe(&mgo.Safe{})

  // fmt.Println("Connected to mongo", db_url)

  // Session = s
  // Mongo = mongo
  Db := Session.DB(db_name)

  return Db, Session
}
