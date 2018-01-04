package request

import (
  // "fmt"
  "os"
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

type Bitaps struct {
  Balance int `json:"balance"`
  TxUnconfirmed int `json:"tx_unconfirmed"`
  TxSent int `json:"tx_sent"`
  TxMultisigReceived int `json:"tx_multisig_received"`
  Pending int `json:"pending"`
  MultisigSent int `json:"multisig_sent"`
  Sent int `json:"sent"`
  TxTotal int `json:"tx_total"`
  TxMultisigSent int `json:"tx_multisig_sent"`
  MultisigReceived int `json:"multisig_received"`
  TxReceived int `json:"tx_received"`
  Received int `json:"received"`
  TxInvalid int `json:"tx_invalid"`
  ConfirmedBalance int `json:"confirmed_balance"`
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

func GetBitapsBalanceWallet() int {
  request := gorequest.New()
  wallet_id := os.Getenv("WALLET_ID")

  url := "https://bitaps.com/api/address/" + wallet_id

  _, body, err := request.Get(url).End()

	if err != nil {
		log.Panic(err)
	}

  res := Bitaps{}
  errs := json.Unmarshal([]byte(body), &res)

  if errs != nil {
    log.Panic(errs)
  }

  return res.Balance
}
