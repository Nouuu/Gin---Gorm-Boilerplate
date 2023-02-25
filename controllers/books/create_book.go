package books

import (
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
)

type createBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pages  uint   `json:"pages" binding:"required"`
}

func (createBookRequest *createBookRequest) ToBook() models.Book {
	return models.Book{
		Title:  createBookRequest.Title,
		Author: createBookRequest.Author,
		Pages:  createBookRequest.Pages,
	}
}
