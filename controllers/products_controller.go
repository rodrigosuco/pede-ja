package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/pede-ja/models"
	"github.com/rodrigosuco/pede-ja/services/products"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func CreateProduct(c *gin.Context)  {
	var product models.Product

	c.ShouldBindBodyWithJSON(&product)

	createdProduct, err := services.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating product, error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": createdProduct,
	})
}
