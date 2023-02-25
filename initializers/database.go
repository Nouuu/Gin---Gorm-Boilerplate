package initializers

import (
	"fmt"
	"github.com/nouuu/gorm-gin-boilerplate/repositories"
	"github.com/nouuu/gorm-gin-boilerplate/repositories/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDatabase() error {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envConf.DbHost,
		envConf.DbUsername,
		envConf.DbPassword,
		envConf.DbName,
		envConf.DbPort,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return err
	}
	repositories.InitDB(db)

	if envConf.DbSync {
		// Comme c'est le dernier appel, on peut retourner directement le résultat de la fonction
		return autoMigrate(db)
	}

	return nil
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.BookEntity{})
	if err != nil {
		return err
	}
	return nil
}
