package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Names      string `json:"names" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Isverified bool
	Cart       Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-To-One relationship (has one)
}
type Cart struct {
	gorm.Model
	UserID   uint `binding:"required"`
	CartItem []CartItem
}
type CartItem struct {
	gorm.Model
	CartID    uint  `binding:"required"`
	Quantity  int32 `binding:"required"`
	ProductID uint  `binding:"required"`
}

type Product struct {
	gorm.Model
	Name  string `json:"name" binding:"required" gorm:"uniqueIndex"`
	Price float32
	Count int32
}
