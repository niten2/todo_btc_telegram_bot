package main

import (
  // "fmt"

  "app-telegram/db"
  "app-telegram/bot_api"
  "app-telegram/logger"
  // "app-telegram/schedule"
)

func main() {
	db.Connect()
  logger.InitFileLogger()

  bot_api.InitBot()
  bot_api.InitActions()
}
