package models

import (
  "fmt"

  // "log"

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
  Value string `json:"value" bson:"value"`
}

func (c *Coin) New (name string, value float64) Coin {
  return Coin{
    ID: bson.NewObjectId(),
    Name: name,
    Value: value,
  }
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

  err := coin_collection.Update(bson.M{"_id": u.ID}, change)

  return err
}

func (c *Coin) FindAll() []Coin {
  var coins []Coin

  coin_collection := db.Db.C("coins")

  err := coin_collection.Find(nil).All(&coins)

  if err != nil {
    return nil, err
  }

  return coins, nil
}

func (c *Coin) FindByName(name string) (Coin, error) {
  coin_collection := db.Db.C("coins")

  coin := Coin{}
  err := coin_collection.Find(bson.M{"name": name}).One(&coin)

  if err != nil {
    return nil, err
  }

  return coin, nil
}

func CreatePoloniexCoinList(coins Coin) string {
  res := "poliniex list \n"

  for _, coin := range coins {
    res = res + fmt.Sprintf("%s %s \n", coin.Name, coin.Value)
  }

  return res
}

func (c *Coin) Fetch() (error) {
  json, err := request.PoloniexRequest()

  if err != nil {
    return err
  }

  for k, v := range json {
    coin, err := Coin.FindByName(v["name"])

    if (err) {
      coin := Coin.New(v["name"], v["last"])
      err := coin.Create()

      if err != nil {
        return err
      }

      return nil
    }

    coin.Value = v["last"]
    coin.Save()

    return nil
  }

  return nil
}



// func CreateCoin(Name string, Value float64) Coin {
//   coin_collection := db.Db.C("coins")

//   coin_document := Coin{Name: Name, Value: Value}

//   coin_collection.Insert(coin_document)

//   return coin_document
// }

// func CreateCoins(Coins []Coin) []Coin {
//   for _, coin := range Coins {
//     CreateCoin(coin.Name, coin.Value)
//   }

//   return Coins
// }

// func BuildCoins(Coins map[string]request.PoloniexCoin) []Coin {
//   var coins []Coin

//   for k, v := range Coins {
//     coin := Coin{Name: k, Value: v.Last}
//     coins = append(coins, coin)
//   }

//   return coins
// }


// func UpdateCoinsPoloniex() {
//   // db, session := db.Connect()
//   // defer session.Close()

//   // CreateCoins(db, BuildCoins(db, request.PoloniexRequest()))
//   log.Printf("coins update")
// }


// func CreatePoloniexAlert(input string) string {
//   values := strings.Split(input, " ")

//   name := fmt.Sprintf("BTC_%s", values[1])
//   compare := values[2]
//   value := values[3]

//   coin := FindCoin(name)

//   // z1, _ := strconv.ParseFloat("0.1", 64)
//   // z2, _  := strconv.ParseFloat("0.02", 64)
//   // fmt.Println(z2 > z1)

//   input_value_int, _ := strconv.ParseFloat(value, 64)
//   coin_value_int, _ := strconv.ParseFloat(coin.Value, 64)


//   fmt.Println(compare)


//   if compare == ">" && input_value_int > coin_value_int {
//     fmt.Println(1111)
//   }

//   if compare == "<" && input_value_int < coin_value_int {
//     fmt.Println(2222)
//   }

//   // fmt.Println(valueInt)


//   // coin.Value

//   // fmt.Println(coin)

//   // fmt.Println(0.0003 > 0.001)

//   // fmt.Println(name, compare, value)


//   return "ok"
// }
