package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nouuu/gorm-gin-boilerplate/controllers"
	"github.com/nouuu/gorm-gin-boilerplate/initializers"
	"log"
	"strings"
)

func main() {
	initializers.Init()

	r := gin.Default()

	controllers.InitRouter(r)

	err := r.Run(strings.Join([]string{":", initializers.Env.Port}, ""))
	if err != nil {
		log.Fatal(err)
	}
}
