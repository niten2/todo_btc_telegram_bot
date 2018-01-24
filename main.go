package main

import (
  // "fmt"

  "app-telegram/db"
  "app-telegram/logger"

  "app-telegram/bot_api"
  "app-telegram/schedule"
)

func main() {
  db.Connect()
  logger.InitFileLogger()

  bot_api.InitBot()
  schedule.InitSchedule()
  bot_api.InitActions()
}
