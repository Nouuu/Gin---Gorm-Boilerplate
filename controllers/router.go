package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers/books"
)

func InitRouter(r *gin.Engine) {
	books.InitBookController(r)
}
