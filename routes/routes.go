package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/pede-ja/controllers"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Application up and running!",
		})
	})

	router.GET("/products", controllers.GetProducts)
}
