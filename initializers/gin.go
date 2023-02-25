package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers"
	"log"
	"strings"
)

var ginEngine *gin.Engine

func initGinEngine() {
	ginEngine = gin.Default()
	controllers.InitRouter(ginEngine)

}

func RunServer() {
	if ginEngine == nil {
		log.Fatal("Gin engine is not initialized")
	}
	err := ginEngine.Run(strings.Join([]string{":", envConf.Port}, ""))
	if err != nil {
		log.Fatal(err)
	}
}
