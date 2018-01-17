package models

import (
  // "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Coin struct {
  ID string `json:"id" bson:"_id,omitempty"`
  Name string `json:"name"`
  Value string `json:"value"`
}

func CreateCoin(Db *mgo.Database, Name string, Value string) Coin {
  coin_collection := Db.C("coins")

  coin_document := Coin{Name: Name, Value: Value}

  coin_collection.Insert(coin_document)

  return coin_document
}

func FindCoin(Db *mgo.Database, Name string) Coin {
  coin_collection := Db.C("coins")

  coin := Coin{}
  err := coin_collection.Find(bson.M{"name": Name}).One(&coin)

  if err != nil {
    panic(err.Error())
  }

  return coin
}
