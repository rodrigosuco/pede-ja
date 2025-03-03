package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/pede-ja/models"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"products": models.Products,
	})
}
