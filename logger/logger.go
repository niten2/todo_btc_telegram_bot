package logger

import (
  "fmt"

  "os"
  "github.com/sirupsen/logrus"

  "app-telegram/config"
)

var Log = logrus.New()

func InitFileLogger() {
  if config.Settings().IsEnvTest {
    Log.Out = os.Stdout
  }

  file, err := os.OpenFile("logrus.log", os.O_CREATE | os.O_WRONLY, 0666)

  if err == nil {
    fmt.Println(err)
    Log.Out = file
  }
}
