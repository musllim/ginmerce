package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"github.com/musllim/ginmerce/controllers"
	_ "github.com/musllim/ginmerce/docs"
	"github.com/musllim/ginmerce/inits"
	"github.com/musllim/ginmerce/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
