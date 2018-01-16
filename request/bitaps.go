package request

import (
  // "fmt"
  "os"
	"log"
  "encoding/json"
  "github.com/parnurzeal/gorequest"
)

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

