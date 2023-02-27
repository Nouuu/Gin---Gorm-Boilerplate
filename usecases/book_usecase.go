package usecases

import (
	"github.com/nouuu/gorm-gin-boilerplate/logger"
	"github.com/nouuu/gorm-gin-boilerplate/repositories"
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"github.com/nouuu/gorm-gin-boilerplate/utils/optional"
)

func GetBook(id uint) (optional.Optional[models.Book], error) {
	logger.DebugPrintln("GetBook(id uint) optional.Optional[models.Book]")
	return repositories.GetBook(id)
}

func GetBooks() []models.Book {
	logger.DebugPrintln("GetBooks() []models.Book")
	return repositories.GetBooks()
}

func CreateBook(book models.Book) (optional.Optional[models.Book], error) {
	logger.DebugPrintln("CreateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.CreateBook(book)
}

func UpdateBook(book models.Book) (optional.Optional[models.Book], error) {
	logger.DebugPrintln("UpdateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.UpdateBook(book)
}

func DeleteBook(id uint) error {
	logger.DebugPrintln("DeleteBook(id uint) error")
	return repositories.DeleteBook(id)
}
