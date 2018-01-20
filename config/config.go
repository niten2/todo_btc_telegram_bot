package config

import (
  "os"
  "strconv"
  "strings"
)

type Setting struct {
  Env string
  DbUrl string
  DbName string
  TelegramToken string
  TelegramUserId int64
  WalletId string
  ScheduleEverySeconds bool
}

func Settings() Setting {
  ScheduleEverySeconds, _ := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

  DbName := strings.Split(os.Getenv("DB_URL"), "/")[3]

  TelegramUserId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)

  return Setting{
    Env: os.Getenv("ENV"),
    DbUrl: os.Getenv("DB_URL"),
    DbName: DbName,
    TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
    TelegramUserId: TelegramUserId,
    WalletId: os.Getenv("WALLET_ID"),
    ScheduleEverySeconds: ScheduleEverySeconds,
  }
}
