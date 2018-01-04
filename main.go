package main

import (
  // "fmt"
  "app-telegram/bot_api"
  "app-telegram/schedule"
  "github.com/joho/godotenv"
)

func main() {
  _ = godotenv.Load()

  bot := bot_api.InitBot()

  schedule.Run(bot)
  bot_api.InitActions(bot)
}
