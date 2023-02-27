package initializers

import (
	"github.com/nouuu/gorm-gin-boilerplate/logger"
	"log"
)

func Init() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}
	err = initLogger()
	if err != nil {
		log.Fatal(err)
	}
	err = connectToDatabase()
	if err != nil {
		logger.Fatal(err)
	}
	initGinEngine()
}
