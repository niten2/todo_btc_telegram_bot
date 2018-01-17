package config

import (
  "os"
  "github.com/joho/godotenv"
  "strconv"
)

type Setting struct {
  DbUrl string
  TelegramToken string
  TelegramUserId string
  WalletId string
  ScheduleEverySeconds bool
}

func init() {
  _ = godotenv.Load()
}

func Settings() Setting {
  ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

  return Setting{
    DbUrl: os.Getenv("DB_URL"),
    TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
    TelegramUserId: os.Getenv("TELEGRAM_USER_ID"),
    WalletId: os.Getenv("WALLET_ID"),
    ScheduleEverySeconds: ScheduleEverySeconds,
  }
}
