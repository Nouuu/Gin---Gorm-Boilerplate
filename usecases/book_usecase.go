package usecases

import (
	"github.com/nouuu/gorm-gin-boilerplate/logs"
	"github.com/nouuu/gorm-gin-boilerplate/repositories"
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"github.com/nouuu/gorm-gin-boilerplate/utils/optional"
)

func GetBook(id uint) (optional.Optional[models.Book], error) {
	logs.DebugLogger.Println("GetBook(id uint) optional.Optional[models.Book]")
	return repositories.GetBook(id)
}

func GetBooks() []models.Book {
	logs.DebugLogger.Println("GetBooks() []models.Book")
	return repositories.GetBooks()
}

func CreateBook(book models.Book) (optional.Optional[models.Book], error) {
	logs.DebugLogger.Println("CreateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.CreateBook(book)
}

func UpdateBook(book models.Book) (optional.Optional[models.Book], error) {
	logs.DebugLogger.Println("UpdateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.UpdateBook(book)
}

func DeleteBook(id uint) error {
	logs.DebugLogger.Println("DeleteBook(id uint) error")
	return repositories.DeleteBook(id)
}
