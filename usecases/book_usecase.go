package usecases

import (
	"github.com/nouuu/gorm-gin-boilerplate/repositories"
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"github.com/nouuu/gorm-gin-boilerplate/utils/optional"
)

func GetBook(id uint) (optional.Optional[models.Book], error) {
	logger.DebugLogger.Println("GetBook(id uint) optional.Optional[models.Book]")
	return repositories.GetBook(id)
}

func GetBooks() []models.Book {
	logger.DebugLogger.Println("GetBooks() []models.Book")
	return repositories.GetBooks()
}

func CreateBook(book models.Book) (optional.Optional[models.Book], error) {
	logger.DebugLogger.Println("CreateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.CreateBook(book)
}

func UpdateBook(book models.Book) (optional.Optional[models.Book], error) {
	logger.DebugLogger.Println("UpdateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.UpdateBook(book)
}

func DeleteBook(id uint) error {
	logger.DebugLogger.Println("DeleteBook(id uint) error")
	return repositories.DeleteBook(id)
}
