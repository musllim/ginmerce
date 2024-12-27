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

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.GET("/profile", middlewares.RequireAuth, controllers.Profile)
	r.Run()
}
