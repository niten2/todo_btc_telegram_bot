package models

import (
  // "fmt"

  // "gopkg.in/mgo.v2"
  // "crypto/sha256"
  "gopkg.in/mgo.v2/bson"
  // "encoding/json"

  "app-telegram/db"
  // "app-telegram/logger"
)

type User struct {
  ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name" bson:"name"`
  IdTelegramm int64 `json:"id_telegramm" bson:"id_telegramm"`
  Alerts []Alert `json:"alerts" bson:"alerts"`
}


func (u *User) Create() error {
  user_collection := db.Db.C("users")
  err := user_collection.Insert(u)

  return err
}

func (u *User) Save() error {
  user_collection := db.Db.C("users")

	change := bson.M{
    "$set": bson.M{
      "name": u.Name,
      "id_telegramm": u.IdTelegramm,
      "alerts": u.Alerts,
    },
  }

  err := user_collection.Update(bson.M{"_id": u.ID}, change)

  return err
}

func NewUser(name string, id_telegramm int64) User {
  return User{
    ID: bson.NewObjectId(),
    Name: name,
    IdTelegramm: id_telegramm,
  }
}

func CreateUser(name string, id_telegramm int64) User {
  user := User{
    ID: bson.NewObjectId(),
    Name: name,
    IdTelegramm: id_telegramm,
  }

  user.Create()

  return user
}

// NOTE find
func FindUserById(id string) (User, error) {
  user_collection := db.Db.C("users")

  user := User{}

  err := user_collection.Find(bson.M{"_id": bson.ObjectId(id)}).One(&user)

  if err != nil {
    return user, err
  }

  return user, nil
}

func FindUserByName(Name string) (User, error) {
  user_collection := db.Db.C("users")

  user := User{}
  err := user_collection.Find(bson.M{"name": Name}).One(&user)

  if err != nil {
    return user, err
  }

  return user, nil
}

func FindByIdTelegramm(id_telegramm int64) (User, error) {
  user_collection := db.Db.C("users")
  user := User{}
  err := user_collection.Find(bson.M{"id_telegramm": id_telegramm}).One(&user)

  if err != nil {
    return user, err
  }

  return user, nil
}
