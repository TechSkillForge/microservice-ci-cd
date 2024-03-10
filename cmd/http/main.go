package main

import (
	"products-msvc/cmd/http/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Define routes
	r.POST("/products", createProductHandler)
	r.GET("/products", getProductsHandler)

	// Start server
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}

var products = []*models.Product{}

type createProductInput struct {
	Name string `json:"name"`
}

func createProductHandler(c *gin.Context) {
	var requestBody createProductInput

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	product, err := models.NewProduct(requestBody.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	products = append(products, product)

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": product})
}

func getProductsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}
