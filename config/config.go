package config

import (
  "fmt"

  "os"
  "strconv"
  "strings"
  "github.com/joho/godotenv"

  // . "app-telegram/logger"
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
    err := godotenv.Load("../.env.test")

    if err != nil {
      err = godotenv.Load(".env.test")

      fmt.Println(111)

      if err != nil {
        panic(err)
      }
    }
  } else {
    err := godotenv.Load()

    if err != nil {
      panic(err)
    }
  }

  if os.Getenv("DB_URL") == "" {
    panic("DB_URL not found")
  }

  DbName := strings.Split(os.Getenv("DB_URL"), "/")[3]

  TelegramUserId, err := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)

  if err != nil {
    fmt.Println(err)
  }

  ScheduleEverySeconds, err := strconv.ParseBool(os.Getenv("SCHEDULE_EVERY_SECONDS"))

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
