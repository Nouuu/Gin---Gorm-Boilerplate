package migrations

import (
	"github.com/nouuu/gorm-gin-boilerplate/repositories/entities"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.BookEntity{})
	if err != nil {
		return err
	}
	return nil
}
