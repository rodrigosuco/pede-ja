package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/pede-ja/routes"
)

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3000")
}
