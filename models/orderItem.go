package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
	Total     float64
}
