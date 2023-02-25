package books

import (
	"github.com/nouuu/gorm-gin-boilerplate/usecases/models"
)

type bookResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  uint   `json:"pages"`
}

func toBookResponse(book models.Book) bookResponse {
	return bookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Pages:  book.Pages,
	}
}

func booksToBookResponses(books []models.Book) []bookResponse {
	var bookResponses = make([]bookResponse, 0)
	for _, book := range books {
		bookResponses = append(bookResponses, toBookResponse(book))
	}
	return bookResponses
}
