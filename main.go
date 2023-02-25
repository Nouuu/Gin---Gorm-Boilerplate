package main

import (
	"github.com/nouuu/gorm-gin-boilerplate/initializers"
)

func main() {
	initializers.Init()

	initializers.RunServer()
}
