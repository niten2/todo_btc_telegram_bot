package main

import (
  // "fmt"
  "github.com/joho/godotenv"

  "app-telegram/bot_api"
  "app-telegram/db"
  "app-telegram/schedule"

  // . "app-telegram/logger"
  // "app-telegram/config"
)

func main() {
  _ = godotenv.Load()
	db.Connect()

  bot := bot_api.InitBot()
  schedule.Run(bot)
  bot_api.InitActions(bot)
}
