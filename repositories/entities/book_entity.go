package entities

import (
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
	"gorm.io/gorm"
)

type BookEntity struct {
	gorm.Model
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
	Pages  uint   `gorm:"column:pages"`
}

func (BookEntity) TableName() string {
	return "books"
}

func ToBook(bookEntity BookEntity) models.Book {
	return models.Book{
		ID:        bookEntity.ID,
		Title:     bookEntity.Title,
		Author:    bookEntity.Author,
		Pages:     bookEntity.Pages,
		CreatedAt: bookEntity.CreatedAt,
		UpdatedAt: bookEntity.UpdatedAt,
	}
}

func BookEntitiesToBooks(bookEntities []BookEntity) []models.Book {
	var books = make([]models.Book, 0)
	for _, bookEntity := range bookEntities {
		books = append(books, ToBook(bookEntity))
	}
	return books
}

func FromBook(book models.Book) BookEntity {
	return BookEntity{
		Model: gorm.Model{
			ID: book.ID,
		},
		Title:  book.Title,
		Author: book.Author,
		Pages:  book.Pages,
	}
}
