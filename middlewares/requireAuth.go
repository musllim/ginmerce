package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
)

func RequireAuth(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil

	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	inits.Db.First(&user, claims["sub"].(string))
	if user.ID == 0 {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
	user.Password = ""

	c.Set("user", user)
	c.Next()
}
