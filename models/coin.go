package models

import (
  "fmt"

  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "app-telegram/request"
  "app-telegram/db"

  // "strings"
  // "strconv"
)

type Coin struct {
  ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name" bson:"name"`
  Value float64 `json:"value" bson:"value"`
}

func (c *Coin) Create() error {
  coin_collection := db.Db.C("coins")

  err := coin_collection.Insert(c)

  return err
}

func (c *Coin) Save() error {
  coin_collection := db.Db.C("coins")

	change := bson.M{
    "$set": bson.M{
      "name": c.Name,
      "value": c.Value,
    },
  }

  err := coin_collection.Update(bson.M{"_id": c.ID}, change)

  return err
}

func FindCoinAll() ([]Coin, error) {
  var coins []Coin

  coin_collection := db.Db.C("coins")

  err := coin_collection.Find(nil).All(&coins)

  if err != nil {
    return nil, err
  }

  return coins, nil
}

func NewCoin (name string, value float64) Coin {
  return Coin{
    ID: bson.NewObjectId(),
    Name: name,
    Value: value,
  }
}

func CreateCoin(name string, value float64) (Coin, error) {
  coin := Coin{
    ID: bson.NewObjectId(),
    Name: name,
    Value: value,
  }

  err := coin.Create()

  return coin, err
}

func FetchCoin() (error) {
  resp, err := request.RequestPoloniex()

  if err != nil {
    return err
  }

  for _, v := range resp {
    coin, err := FindCoinByName(v.Name)

    if err != nil {
      coin := NewCoin(v.Name, v.Value)

      err := coin.Create()

      if err != nil {
        return err
      }

    } else {
      coin.Value = v.Value
      coin.Save()
    }
  }

  return nil
}

// NOTE find
func FindCoinById(id string) (Coin, error) {
  coin_collection := db.Db.C("coins")

  coin := Coin{}

  err := coin_collection.Find(bson.M{"_id": bson.ObjectId(id)}).One(&coin)

  return coin, err
}

func FindCoinByName(name string) (Coin, error) {
  coin_collection := db.Db.C("coins")

  coin := Coin{}
  err := coin_collection.Find(bson.M{"name": name}).One(&coin)

  if err != nil {
    return coin, err
  }

  return coin, nil
}

// Actions
func CreatePoloniexCoinList(coins []Coin) string {
  res := "Poliniex list \n"

  for _, coin := range coins {
    res = res + fmt.Sprintf("%s %f \n", coin.Name, coin.Value)
  }

  return res
}
