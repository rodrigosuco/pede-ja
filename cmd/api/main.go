package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println("Hello World!")

	r := gin.Default()
  r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Application up and runnning!",
    })
  })
  r.Run(":3000")
}