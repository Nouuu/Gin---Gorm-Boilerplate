package repositories

import (
	"github.com/nouuu/gorm-gin-boilerplate/repositories/entities"
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"github.com/nouuu/gorm-gin-boilerplate/utils/optional"
)

func GetBooks() []models.Book {
	var bookEntities []entities.BookEntity
	db.Find(&bookEntities)

	return entities.BookEntitiesToBooks(bookEntities)
}

func GetBook(id uint) (optional.Optional[models.Book], error) {
	var bookEntity entities.BookEntity
	err := db.First(&bookEntity, id).Error
	if err != nil {
		return optional.Empty[models.Book](), err
	}
	return optional.Of[models.Book](entities.ToBook(bookEntity)), nil
}

func CreateBook(book models.Book) (optional.Optional[models.Book], error) {
	bookEntity := entities.FromBook(book)
	err := db.Create(&bookEntity).Error
	if err != nil {
		return optional.Empty[models.Book](), err
	}
	return optional.Of[models.Book](entities.ToBook(bookEntity)), nil
}

func UpdateBook(book models.Book) (optional.Optional[models.Book], error) {
	bookEntity := entities.FromBook(book)
	err := db.Model(&bookEntity).Where("id = ?", book.ID).Updates(bookEntity).Error
	if err != nil {
		return optional.Empty[models.Book](), err
	}
	return optional.Of[models.Book](entities.ToBook(bookEntity)), nil
}

func DeleteBook(id uint) error {
	err := db.Delete(&entities.BookEntity{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
