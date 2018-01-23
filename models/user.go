package models

import (
  // "fmt"

  "gopkg.in/mgo.v2/bson"

  // "app-telegram/logger"
  "app-telegram/db"
)

type User struct {
  ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name" bson:"name"`
  IdTelegram int64 `json:"id_telegram" bson:"id_telegram"`
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
      "id_telegram": u.IdTelegram,
      "alerts": u.Alerts,
    },
  }

  err := user_collection.Update(bson.M{"_id": u.ID}, change)

  return err
}

func NewUser(name string, id_telegram int64) User {
  return User{
    ID: bson.NewObjectId(),
    Name: name,
    IdTelegram: id_telegram,
  }
}

func CreateUser(name string, id_telegram int64) (User, error) {
  user := User{
    ID: bson.NewObjectId(),
    Name: name,
    IdTelegram: id_telegram,
  }

  err := user.Create()

  return user, err
}

// NOTE find
func FindUserAll() ([]User, error) {
  var users []User

  user_collection := db.Db.C("users")

  err := user_collection.Find(nil).All(&users)

  return users, err
}

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

func FindUserByIdTelegram(id_telegram int64) (User, error) {
  user_collection := db.Db.C("users")
  user := User{}
  err := user_collection.Find(bson.M{"id_telegram": id_telegram}).One(&user)

  if err != nil {
    return user, err
  }

  return user, nil
}

func FindOrCreateUserByIdTelegram(user_name string, id_telegram int64) (User, error) {
  user, err := FindUserByIdTelegram(id_telegram)

  if err != nil && err.Error() == "not found" {
    user, err = CreateUser(user_name, id_telegram)
  }

  return user, err
}
