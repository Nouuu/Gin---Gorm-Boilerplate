package books

import (
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
)

type updateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pages  uint   `json:"pages" binding:"required"`
}

func (updateBookRequest *updateBookRequest) ToBook(id uint) models.Book {
	return models.Book{
		ID:     id,
		Title:  updateBookRequest.Title,
		Author: updateBookRequest.Author,
		Pages:  updateBookRequest.Pages,
	}
}
