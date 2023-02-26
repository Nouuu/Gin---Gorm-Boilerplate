package initializers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers"
	"github.com/nouuu/gorm-gin-boilerplate/logs"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var ginEngine *gin.Engine

func initGinEngine() {
	setGinMode()
	ginEngine = gin.New()
	ginEngine.Use(gin.Recovery())
	setLogger()
	controllers.InitRouter(ginEngine)
}

func setGinMode() {
	if envConf.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func setLogger() {
	f, err := os.OpenFile(envConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	ginEngine.Use(
		gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			if param.StatusCode < 400 && envConf.ReleaseMode {
				return ""
			}
			return fmt.Sprintf("[%s][GIN] %s - \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format(time.RFC3339),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
}

func RunServer() {
	if ginEngine == nil {
		logs.ErrorLogger.Fatal("Gin engine is not initialized")
	}
	err := ginEngine.Run(strings.Join([]string{":", envConf.Port}, ""))
	if err != nil {
		logs.ErrorLogger.Fatal(err)
	}
}
