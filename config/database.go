package config

import (
	"log"

	"github.com/Ferdiansyah4343/e-commerce-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "devdb:devdb123@tcp(127.0.0.1:3306)/go_commerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	DB = db
	migrateDB()
}

func migrateDB() {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Order{},
		&models.OrderItem{},
	}
	if err := DB.AutoMigrate(modelsToMigrate...); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
