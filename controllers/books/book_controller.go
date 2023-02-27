package books

import (
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers"
	"github.com/nouuu/gorm-gin-boilerplate/usecases"
	"net/http"
	"strconv"
)

func InitBookController(r *gin.Engine) {
	routerGroup := r.Group("/books")
	routerGroup.GET("", getBooks)
	routerGroup.GET("/:id", getBook)
	routerGroup.POST("", addBook)
	routerGroup.PUT("/:id", updateBook)
	routerGroup.DELETE("/:id", deleteBook)
}

func getBooks(c *gin.Context) {
	books := usecases.GetBooks()
	c.JSON(http.StatusOK, booksToBookResponses(books))
}

func getBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		controllers.HandleError(c, http.StatusBadRequest, err)
		return
	}
	optionalBook, err := usecases.GetBook(uint(id))
	if err != nil {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	if optionalBook.IsEmpty() {
		controllers.HandleError(c, http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, toBookResponse(optionalBook.Get()))
}

func addBook(c *gin.Context) {
	var createBookRequest createBookRequest
	if err := c.ShouldBindJSON(&createBookRequest); err != nil {
		controllers.HandleError(c, http.StatusBadRequest, err)
		return
	}
	optionalBook, err := usecases.CreateBook(createBookRequest.ToBook())
	if err != nil {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	if optionalBook.IsEmpty() {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, toBookResponse(optionalBook.Get()))
}

func updateBook(c *gin.Context) {
	var updateBookRequest updateBookRequest
	if err := c.ShouldBindJSON(&updateBookRequest); err != nil {
		controllers.HandleError(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		controllers.HandleError(c, http.StatusBadRequest, err)
		return
	}
	optionalBook, err := usecases.UpdateBook(updateBookRequest.ToBook(uint(id)))
	if err != nil {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	if optionalBook.IsEmpty() {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, toBookResponse(optionalBook.Get()))
}

func deleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		controllers.HandleError(c, http.StatusBadRequest, err)
		return
	}
	err = usecases.DeleteBook(uint(id))
	if err != nil {
		controllers.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
