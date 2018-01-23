package request

import (
  // "fmt"

	// "log"
  // "os"
  "net/http"
  "encoding/json"
  "io/ioutil"

  // "app-telegram/models"

  "strconv"
)

type CoinPoloniex struct {
  Name string
  Value float64 `json:"last"`
}

func RequestPoloniex() ([]CoinPoloniex, error) {
  url_poloniex := "https://poloniex.com/public?command=returnTicker"

	var coins map[string]interface{}
  var result []CoinPoloniex

  res, err := http.Get(url_poloniex)

	if err != nil {
    return nil, err
	}

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)

	if err != nil {
    return nil, err
	}

	json.Unmarshal(body, &coins)

  for v, k := range coins {
    last := k.(map[string]interface{})["last"].(string)
    value, err := strconv.ParseFloat(last, 64)

    if err != nil {
      return nil, err
    }

    result = append(result, CoinPoloniex{
      Name: v,
      Value: value,
    })
  }

  return result, nil
}

