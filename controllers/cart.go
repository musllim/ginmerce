package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

func GetCart(c *gin.Context) {
	var cart models.Cart
	inits.Db.Where("user_id = ?", c.MustGet("user").(models.User).ID).First(&cart)

	if cart.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Cart not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": cart,
	})
}

func CreateCart(c *gin.Context) {
	cart := models.Cart{UserID: c.MustGet("user").(models.User).ID}
	inits.Db.Create(&cart)

	if cart.ID == 0 {
		c.JSON(500, gin.H{
			"message": "Failed to create cart",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": cart,
	})
}

func CreateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	c.BindJSON(&cartItem)
	inits.Db.Create(&cartItem)

	if cartItem.ID == 0 {
		c.JSON(500, gin.H{
			"message": "Failed to create cart item",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": cartItem,
	})
}

func GetCartItems(c *gin.Context) {
	var cartItems []models.CartItem
	inits.Db.Model(&models.CartItem{}).Where("cart_id = ?", c.Param("id")).Find(&cartItems)

	if len(cartItems) == 0 {
		c.JSON(404, gin.H{
			"message": "Cart items not found, cart is empty",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": cartItems,
	})
}
