package initializers

import (
	"github.com/nouuu/gorm-gin-boilerplate/logger"
	"io"
	"os"
)

func initLogger() error {
	f, err := os.OpenFile(envConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	loggerOutput := io.MultiWriter(f, os.Stdout)
	infoEnabled := envConf.LogLevel == "info" || envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
	warningEnabled := envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
	debugEnabled := envConf.LogLevel == "debug"

	logger.InitLoggers(loggerOutput, infoEnabled, warningEnabled, debugEnabled)
	return nil
}
