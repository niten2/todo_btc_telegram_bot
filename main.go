package main

import (
  // "fmt"
  "app-telegram/bot"
  "app-telegram/schedule"
  "app-telegram/action"
  // "app-telegram/request"
)

func main() {
  bot := bot.InitBot()

  schedule.Run(bot)
  action.Run(bot)

}
