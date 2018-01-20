package logger

import (
  "os"
  "github.com/sirupsen/logrus"
  "app-telegram/config"
)

var Log = logrus.New()

func init() {
  env := config.Settings().Env

  if (env == "test") {
    Log.Out = os.Stdout
    return
  }

  file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)

  if err == nil {
    Log.Out = file
  }
}
