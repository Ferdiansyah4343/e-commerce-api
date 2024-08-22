package controllers

import (
	"net/http"

	"github.com/Ferdiansyah4343/e-commerce-api/config"
	"github.com/Ferdiansyah4343/e-commerce-api/helpers"
	"github.com/Ferdiansyah4343/e-commerce-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		helpers.ResponseJson(c.Writer, "Get all Products failed", err.Error(), http.StatusInternalServerError, nil)
		return
	}
	helpers.ResponseJson(c.Writer, "Get all products successfully", "success", http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		helpers.ResponseJson(c.Writer, "Product not found", err.Error(), http.StatusNotFound, nil)
		return
	}
	helpers.ResponseJson(c.Writer, "Get product successfully", "success", http.StatusOK, product)
}

func PostProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		helpers.ResponseJson(c.Writer, "Invalid request payload", err.Error(), http.StatusBadRequest, nil)
		return
	}
	if err := config.DB.Create(&product).Error; err != nil {
		helpers.ResponseJson(c.Writer, "Create product failed", err.Error(), http.StatusInternalServerError, nil)
		return
	}
	helpers.ResponseJson(c.Writer, "Create product successfully", "success", http.StatusCreated, product)
}

func EditProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).Error; err != nil {
		helpers.ResponseJson(c.Writer, "Product not found", err.Error(), http.StatusNotFound, nil)
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		helpers.ResponseJson(c.Writer, "Invalid request payload", err.Error(), http.StatusBadRequest, nil)
		return
	}

	if err := config.DB.Model(&product).Updates(updateData).Error; err != nil {
		helpers.ResponseJson(c.Writer, "Invalid request payload", err.Error(), http.StatusInternalServerError, nil)
		return
	}
	helpers.ResponseJson(c.Writer, "Update product successfully", "success", http.StatusOK, product)
}
