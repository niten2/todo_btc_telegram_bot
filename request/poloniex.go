package request

import (
  "fmt"
	// "log"
  "net/http"
  "encoding/json"
  "io/ioutil"
  // "app-telegram/models"

  // "os"
)

type PoloniexCoin struct {
  Name string
  Last string `json:"last"`
}

func PoloniexRequest() (map[string]PoloniexCoin, error) {
  url_poloniex := "https://poloniex.com/public?command=returnTicker"

  res, err := http.Get(url_poloniex)

	if err != nil {
    fmt.Println("error", err)
	}

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)

	if err != nil {
    fmt.Println("error", err)
    return nil, err
	}

	var coins map[string]PoloniexCoin

	json.Unmarshal(body, &coins)

  return coins, nil
}
