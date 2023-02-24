package books

import (
	"github.com/gin-gonic/gin"
	repository "github.com/nouuu/gorm-gin-boilerplate/repositories"
	"strconv"
)

func Init(r *gin.Engine) {
	routerGroup := r.Group("/books")
	getBooks(routerGroup)
	getBook(routerGroup)
	addBook(routerGroup)
	updateBook(routerGroup)
	deleteBook(routerGroup)
}

func getBooks(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		books := repository.GetBooks()
		c.JSON(200, books)
	})
}

func getBook(r *gin.RouterGroup) {
	r.GET("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		book, err := repository.GetBook(uint(id))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, book)
	})
}

func addBook(r *gin.RouterGroup) {
	r.POST("", func(c *gin.Context) {
		var createBook CreateBookRequest
		if err := c.ShouldBindJSON(&createBook); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		book, err := repository.CreateBook(createBook.ToBook())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, book)
	})
}

func updateBook(r *gin.RouterGroup) {
	r.PUT("/:id", func(c *gin.Context) {
		var updateBook UpdateBookRequest
		if err := c.ShouldBindJSON(&updateBook); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		book, err := repository.UpdateBook(updateBook.ToBook(uint(id)))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, book)
	})
}

func deleteBook(r *gin.RouterGroup) {
	r.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = repository.DeleteBook(uint(id))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Book deleted"})
	})
}
