package repository

import (
	"github.com/nouuu/gorm-gin-boilerplate/initializers"
	"github.com/nouuu/gorm-gin-boilerplate/models"
	"github.com/nouuu/gorm-gin-boilerplate/repositories/entities"
)

func GetBooks() []models.Book {
	var bookEntities []entities.BookEntity
	initializers.DB.Find(&bookEntities)

	return entities.BookEntitiesToBooks(bookEntities)
}

func GetBook(id uint) (models.Book, error) {
	var bookEntity entities.BookEntity
	err := initializers.DB.First(&bookEntity, id).Error
	if err != nil {
		return models.Book{}, err
	}
	return entities.ToBook(bookEntity), nil
}

func CreateBook(book models.Book) (models.Book, error) {
	bookEntity := entities.FromBook(book)
	err := initializers.DB.Create(&bookEntity).Error
	if err != nil {
		return models.Book{}, err
	}
	return entities.ToBook(bookEntity), nil
}

func UpdateBook(book models.Book) (models.Book, error) {
	bookEntity := entities.FromBook(book)
	err := initializers.DB.Model(&bookEntity).Where("id = ?", book.ID).Updates(bookEntity).Error
	if err != nil {
		return models.Book{}, err
	}
	return entities.ToBook(bookEntity), nil
}

func DeleteBook(id uint) error {
	err := initializers.DB.Delete(&entities.BookEntity{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
