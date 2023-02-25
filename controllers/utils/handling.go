package utils

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
