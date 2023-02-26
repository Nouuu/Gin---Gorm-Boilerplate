package initializers

import (
	"github.com/nouuu/gorm-gin-boilerplate/logs"
	"log"
)

func Init() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}
	initLogger()
	err = connectToDatabase()
	if err != nil {
		logs.ErrorLogger.Fatal(err)
	}
	initGinEngine()
}
