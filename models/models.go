package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Names      string `json:"names" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Isverified bool
	Cart       Cart
}
type Cart struct {
	gorm.Model
	UserID   uint
	CartItem []CartItem
}
type CartItem struct {
	gorm.Model
	CartID    uint
	Quantity  int32
	ProductID uint
}

type Product struct {
	gorm.Model
	Name  string `json:"name" binding:"required" gorm:"uniqueIndex"`
	Price float32
	Count int32
}
