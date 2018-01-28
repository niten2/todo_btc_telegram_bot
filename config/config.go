package config

import (
	// "fmt"

	"github.com/joho/godotenv"
	"os"
	// "strconv"
	"strings"
)

type Setting struct {
	Env                  string
	IsEnvTest            bool
	DbUrl                string
	DbName               string
	TelegramToken        string
	ScheduleEverySeconds bool
}

func Settings() Setting {
	if os.Getenv("ENV") == "test" {
		err := godotenv.Load("../.env.test")

		if err != nil {
			err = godotenv.Load(".env.test")

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

	return Setting{
		Env:           os.Getenv("ENV"),
		IsEnvTest:     os.Getenv("ENV") == "test",
		DbUrl:         os.Getenv("DB_URL"),
		DbName:        DbName,
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}
}
