package request

import (
  // "fmt"

  "log"
  "encoding/json"
  "github.com/parnurzeal/gorequest"
)

type Coinmarketcap struct {
  Id string `json:"id"`
  Name string `json:"name"`
  PriceUsd string `json:"price_usd"`
  LastUpdated string `json:"last_updated"`
}

func GetCoinmarketcapCurrentBtc() string {
  request := gorequest.New()
  url := "https://api.coinmarketcap.com/v1/ticker/bitcoin/"

  _, body, err := request.Get(url).End()

  if err != nil {
    log.Panic(err)
  }

  res := []Coinmarketcap{}
  errs := json.Unmarshal([]byte(body), &res)

  if errs != nil {
    log.Panic(errs)
  }

  return res[0].PriceUsd
}
