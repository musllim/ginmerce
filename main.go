package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"github.com/musllim/ginmerce/controllers"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/middlewares"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	inits.ConectDb()
	inits.Migrate()
}
func main() {
	r := gin.Default()
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", middlewares.RequireAuth, controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProduct)
	r.DELETE("/products/:id", middlewares.RequireAuth, controllers.DeleteProduct)

	r.POST("/cart", middlewares.RequireAuth, controllers.CreateCart)
	r.GET("/cart", middlewares.RequireAuth, controllers.GetCart)
	r.POST("/cart/items", middlewares.RequireAuth, controllers.CreateCartItem)
	r.GET("/cart/:id/items", middlewares.RequireAuth, controllers.GetCartItems)

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.GET("/profile", middlewares.RequireAuth, controllers.Profile)
	r.GET("/logout", middlewares.RequireAuth, controllers.Logout)
	r.Run()
}
