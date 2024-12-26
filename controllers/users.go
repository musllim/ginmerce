package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)

	inits.Db.Create(&user)
	user.Password = ""

	c.JSON(200, gin.H{
		"data": user,
	})
}
func Login(c *gin.Context) {
	type User struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var user User
	c.BindJSON(&user)

	var users models.User
	inits.Db.Where("email = ?", user.Email).First(&users)
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid creadentials",
			"error":   err.Error(),
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(users.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid creadentials",
			"error":   err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
