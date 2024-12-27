package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

func CreateCart(c *gin.Context) {
	cart := models.Cart{UserID: c.MustGet("user").(models.User).ID}
	inits.Db.Create(&cart)
	c.JSON(200, gin.H{
		"data": cart,
	})
}
