package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

type CartItem struct {
	CartID    uint  `binding:"required"`
	Quantity  int32 `binding:"required"`
	ProductID uint  `binding:"required"`
}

type Cart struct {
	UserID   uint `binding:"required"`
	CartItem []CartItem
}

// GetCart godoc
// @Summary Get cart
// @Description Get cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} Cart
// @Router /cart [get]
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

// CreateCart godoc
// @Summary Create cart
// @Description Create cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} Cart
// @Router /cart [post]
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

// CreateCartItem godoc
// @Summary Create cart item
// @Description Create cart item
// @Tags cart
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param cartItem body CartItem true "Cart Item"
// @Success 200 {object} CartItem
// @Router /cart/items [post]
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

// GetCartItems godoc
// @Summary Get cart items
// @Description Get cart items
// @Tags cart
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} CartItem
// @Router /cart/{id}/items [get]
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
