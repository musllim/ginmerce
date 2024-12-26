package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Userid int32
}

type CartItem struct {
	gorm.Model
	Cartid    int32
	Productid int32
	Quantity  int32
}

type Product struct {
	gorm.Model
	Name  string
	Price float32
	Count int32
}

type User struct {
	gorm.Model
	Email      string
	Names      string
	Password   string
	Isverified bool
}
