package books

import (
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers/utils"
	repository "github.com/nouuu/gorm-gin-boilerplate/repositories"
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
	books := repository.GetBooks()
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	book, err := repository.GetBook(uint(id))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func addBook(c *gin.Context) {
	var createBookRequest CreateBookRequest
	if err := c.ShouldBindJSON(&createBookRequest); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	book, err := repository.CreateBook(createBookRequest.ToBook())
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) {
	var updateBookRequest UpdateBookRequest
	if err := c.ShouldBindJSON(&updateBookRequest); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	book, err := repository.UpdateBook(updateBookRequest.ToBook(uint(id)))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	err = repository.DeleteBook(uint(id))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
