package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Names      string `json:"names" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Isverified bool
}
type Cart struct {
	gorm.Model
	UserID   uint `binding:"required"`
	User     User
	CartItem []CartItem
}
type CartItem struct {
	gorm.Model
	CartID    uint  `binding:"required"`
	Quantity  int32 `binding:"required" gorm:"check:quantity > 0"`
	ProductID uint  `binding:"required"`
}

type Product struct {
	gorm.Model
	Name  string  `json:"name" binding:"required" gorm:"uniqueIndex"`
	Price float32 `json:"price" binding:"required" gorm:"check:price > 0"`
	Count int32   `json:"count" binding:"required" gorm:"check:count > 0"`
}
