package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	inits.Db.Model(&models.Product{}).Find(&products)
	c.JSON(200, gin.H{
		"data": products,
	})
}

func GetProduct(c *gin.Context) {
	var product models.Product
	inits.Db.First(&product, c.Param("id"))
	c.JSON(200, gin.H{
		"data": product,
	})
}

func DeleteProduct(c *gin.Context) {
	inits.Db.Delete(&models.Product{}, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Product deleted",
	})
}
func CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	inits.Db.Create(&product)
	c.JSON(200, gin.H{
		"data": product,
	})
}
