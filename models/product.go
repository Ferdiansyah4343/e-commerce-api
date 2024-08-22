package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int        `json:"stock"`
	Categories  []Category `gorm:"many2many:product_categories"`
}
