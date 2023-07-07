package main

import (
	"github.com/gin-gonic/gin"
	"goTest04-dbConnection/controllers/productController"
	"goTest04-dbConnection/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/products", productController.GetProducts)
	r.GET("/api/products/:id", productController.GetProductById)
	r.POST("/api/product", productController.InsertProduct)
	r.PUT("/api/product/:id", productController.UpdateProduct)
	r.DELETE("/api/product", productController.DeleteProduct)
	r.Run("localhost:9090")
}
