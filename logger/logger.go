package logger

import (
  "fmt"
  "app-telegram/config"

  "os"
  "github.com/sirupsen/logrus"
)

var Log = logrus.New()

func getLogger() {
  env := config.Settings().Env

  if (env == "test") {
    Log.Out = os.Stdout
    return
  } else {
    file, err := os.OpenFile("logrus.log", os.O_CREATE | os.O_WRONLY, 0666)

    if err == nil {
      fmt.Println(err)
      Log.Out = file
    }
  }
}
