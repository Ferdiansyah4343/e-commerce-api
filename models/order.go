package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Total  float64
	Status string
	Items  []OrderItem
}
