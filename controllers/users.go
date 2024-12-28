package controllers

import (
	"fmt"
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

type User struct {
	Email    string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Names    string `json:"names" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /register [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)

	vvv := inits.Db.Create(&user)
	fmt.Println(vvv.Error.Error())
	user.Password = ""
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 200 {string} string	"token"
// @Router /login [post]
func Login(c *gin.Context) {
	type User struct {
		Email    string `json:"email" binding:"required" email:"true"`
		Password string `json:"password" binding:"required"`
	}
	var user User
	c.BindJSON(&user)

	var users models.User
	inits.Db.Where("email = ?", user.Email).First(&users)

	if users.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid creadentials",
		})
		return
	}

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

// Profile godoc
// @Summary Profile
// @Description Profile
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} User
// @Router /profile [get]
func Profile(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"data": user,
	})
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string	"Logged out"
// @Router /logout [get]
func Logout(c *gin.Context) {
	c.Set("user", nil)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(200, gin.H{
		"message": "Logged out",
	})
}
