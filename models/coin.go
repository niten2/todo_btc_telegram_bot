package models

import (
  // "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "app-telegram/request"
  "app-telegram/db"
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

func CreateCoins(Db *mgo.Database, Coins []Coin) []Coin {
  for _, coin := range Coins {
    CreateCoin(Db, coin.Name, coin.Value)
  }

  return Coins
}

func BuildCoins(Db *mgo.Database, Coins map[string]request.PoloniexCoin) []Coin {
  var coins []Coin

  for k, v := range Coins {
    coin := Coin{Name: k, Value: v.Last}
    coins = append(coins, coin)
  }

  return coins
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

func UpdateCoinsPoloniex() {
  db, session := db.Connect()
  defer session.Close()

  CreateCoins(db, BuildCoins(db, request.PoloniexRequest()))
  log.Printf("coins update")
}
