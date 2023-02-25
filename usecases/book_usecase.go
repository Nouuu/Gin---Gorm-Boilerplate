package usecases

import (
	"github.com/nouuu/gorm-gin-boilerplate/repositories"
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"github.com/nouuu/gorm-gin-boilerplate/utils/optional"
	"log"
)

func GetBook(id uint) (optional.Optional[models.Book], error) {
	log.Println("GetBook(id uint) optional.Optional[models.Book]")
	return repositories.GetBook(id)
}

func GetBooks() []models.Book {
	log.Println("GetBooks() []models.Book")
	return repositories.GetBooks()
}

func CreateBook(book models.Book) (optional.Optional[models.Book], error) {
	log.Println("CreateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.CreateBook(book)
}

func UpdateBook(book models.Book) (optional.Optional[models.Book], error) {
	log.Println("UpdateBook(book models.Book) optional.Optional[models.Book]")
	return repositories.UpdateBook(book)
}

func DeleteBook(id uint) error {
	log.Println("DeleteBook(id uint) error")
	return repositories.DeleteBook(id)
}
