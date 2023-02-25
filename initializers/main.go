package initializers

import "log"

func Init() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}
	err = connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	initGinEngine()
}
