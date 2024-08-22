package test

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	dsn := "devdb:devdb123@tcp(127.0.0.1:3306)/go_commerce?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
}
