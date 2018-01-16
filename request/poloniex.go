package request

import (
  _ "fmt"
  // "os"
	"log"
  "net/http"
  "encoding/json"
  "io/ioutil"
  // "github.com/parnurzeal/gorequest"
)


type PoloniexCoin struct {
  Last string `json:"last"`
}

func PoloniexRequest() map[string]PoloniexCoin {
  url_poloniex := "https://poloniex.com/public?command=returnTicker"

  res, err := http.Get(url_poloniex)

	if err != nil {
		log.Panic(err)
	}

  body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Panic(err)
	}

	var coins map[string]PoloniexCoin

	json.Unmarshal(body, &coins)

  return coins
}
