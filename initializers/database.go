package initializers

import (
	"fmt"
	"github.com/nouuu/gorm-gin-boilerplate/repositories/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDatabase() error {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Env.DbHost,
		Env.DbUsername,
		Env.DbPassword,
		Env.DbName,
		Env.DbPort,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return err
	}
	DB = db

	if Env.DbSync {
		// Comme c'est le dernier appel, on peut retourner directement le résultat de la fonction
		return autoMigrate()
	}

	return nil
}

func autoMigrate() error {
	err := DB.AutoMigrate(&entities.BookEntity{})
	if err != nil {
		return err
	}
	return nil
}
