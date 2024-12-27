package inits

import "github.com/musllim/ginmerce/models"

func Migrate() {
	Db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{})
}
