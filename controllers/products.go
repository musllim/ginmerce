package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

type Product struct {
	Name  string `json:"name" binding:"required" gorm:"uniqueIndex"`
	Price float32
	Count int32
}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	inits.Db.Model(&models.Product{}).Find(&products)
	c.JSON(200, gin.H{
		"data": products,
	})
}

// GetProduct godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} Product
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	var product models.Product
	inits.Db.First(&product, c.Param("id"))
	c.JSON(200, gin.H{
		"data": product,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {string} string	"Product deleted"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	inits.Db.Delete(&models.Product{}, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "Product deleted",
	})
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body Product true "Product"
// @Success 200 {object} Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	inits.Db.Create(&product)
	c.JSON(200, gin.H{
		"data": product,
	})
}
