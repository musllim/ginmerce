package inits

import "github.com/musllim/ginmerce/models"

func Migrate() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Product{})
	Db.AutoMigrate(&models.Cart{})
	Db.AutoMigrate(&models.CartItem{})
}
