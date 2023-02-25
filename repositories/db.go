package repositories

import "gorm.io/gorm"

var db *gorm.DB

func InitDB(database *gorm.DB) {
	if database == nil {
		panic("database is nil")
	}
	db = database
}
