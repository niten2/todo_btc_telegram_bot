package config

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "github.com/joho/godotenv"
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
  if os.Getenv("ENV") == "test" {
    _ = godotenv.Load("../.env.test")
  } else {
    _ = godotenv.Load("../.env")
  }

  ScheduleEverySeconds, err := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

  if err != nil {
    fmt.Println(err)
  }

  var DbName string
  if os.Getenv("DB_URL") == "" {
    fmt.Println("env DB_URL not found")
  } else {
    DbName = strings.Split(os.Getenv("DB_URL"), "/")[3]
  }

  TelegramUserId, err := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)

  if err != nil {
    fmt.Println(err)
  }

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
