package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

func GetCart(c *gin.Context) {
	var cart models.Cart
	inits.Db.Where("user_id = ?", c.MustGet("user").(models.User).ID).First(&cart)
	c.JSON(200, gin.H{
		"data": cart,
	})
}

func CreateCart(c *gin.Context) {
	cart := models.Cart{UserID: c.MustGet("user").(models.User).ID}
	inits.Db.Create(&cart)
	c.JSON(200, gin.H{
		"data": cart,
	})
}

func CreateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	c.BindJSON(&cartItem)
	inits.Db.Create(&cartItem)
	c.JSON(200, gin.H{
		"data": cartItem,
	})
}

func GetCartItems(c *gin.Context) {
	var cartItems []models.CartItem
	inits.Db.Model(&models.CartItem{}).Where("cart_id = ?", c.Param("id")).Find(&cartItems)
	c.JSON(200, gin.H{
		"data": cartItems,
	})
}
