package books

import "github.com/nouuu/gorm-gin-boilerplate/models"

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pages  uint   `json:"pages" binding:"required"`
}

func (createBookRequest *CreateBookRequest) ToBook() models.Book {
	return models.Book{
		Title:  createBookRequest.Title,
		Author: createBookRequest.Author,
		Pages:  createBookRequest.Pages,
	}
}
