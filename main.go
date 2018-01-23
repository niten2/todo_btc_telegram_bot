package main

import (
  // "fmt"

  "app-telegram/db"
  "app-telegram/bot_api"
  // "app-telegram/schedule"
)

func main() {
	db.Connect()

  bot_api.InitBot()
  bot_api.InitActions()
}
