package initializers

import (
	"github.com/nouuu/gorm-gin-boilerplate/logs"
	"io"
	"log"
	"os"
)

func initLogger() {
	f, err := os.OpenFile(envConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	loggerOutput := io.MultiWriter(f, os.Stdout)
	infoEnabled := envConf.LogLevel == "info" || envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
	warningEnabled := envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
	debugEnabled := envConf.LogLevel == "debug"

	logs.InitLoggers(loggerOutput, infoEnabled, warningEnabled, debugEnabled)
}
