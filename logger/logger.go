package logger

import (
	// "fmt"

	"github.com/sirupsen/logrus"
	"os"
	// "app-telegram/config"
)

var Log = logrus.New()

func InitFileLogger() {
	Log.Out = os.Stdout

	// if config.Settings().IsEnvTest {
	//   Log.Out = os.Stdout
	// }

	// file, err := os.OpenFile("logrus.log", os.O_CREATE | os.O_WRONLY, 0666)

	// if err == nil {
	//   fmt.Println(err)
	//   Log.Out = file
	// }
}
