package initializers

import (
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
		logger.ErrorLogger.Fatal(err)
	}
	initGinEngine()
}
