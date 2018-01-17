package request

import (
  // "fmt"
	"log"
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

func PoloniexRequest() map[string]PoloniexCoin {

  // models.Coin

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
	// var coins map[string]models.Coin
	// var coins map[string]string
  // fmt.Println(body)


  // fmt.Println(models.Coin{})

  // fmt.Println(body)

	json.Unmarshal(body, &coins)

  // fmt.Println(coins)




  return coins
}
