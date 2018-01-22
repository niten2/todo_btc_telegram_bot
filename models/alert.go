package models

import (
  "fmt"
  // "gopkg.in/mgo.v2"
  // "crypto/sha256"
  // "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2/bson"
  // "encoding/json"

  "strings"
  "strconv"

  // "app-telegram/db"
  // "app-telegram/types"
)

type Alert struct {
  ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name"`
  Compare string `json:"compare"`
  Value float64 `json:"value"`
}

func (a *Alert) AsMap() map[string]string {
  alert := make(map[string]string)

  alert["id"] = string(a.ID)
  alert["name"] = a.Name
  alert["value"] = fmt.Sprintf("%v", a.Value)

  return alert
}

// example input "p SBD > 0.000020"
func NewAlert(input string) (Alert, error) {
  values := strings.Split(input, " ")

  name := fmt.Sprintf("BTC_%s", values[1])
  compare := values[2]
  value, err := strconv.ParseFloat(values[3], 64)

  if err != nil {
    return Alert{}, err
  }

  return Alert{
    ID: bson.NewObjectId(),
    Name: name,
    Compare: compare,
    Value: value,
  }, nil
}

