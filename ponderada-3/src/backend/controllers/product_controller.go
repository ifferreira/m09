package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend/config"
	"backend/models"
	"backend/utils"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController() *ProductController {
	return &ProductController{
		DB: config.GetDB(),
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product.ToResponse()})
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := pc.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	var productResponses []models.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, product.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{"products": productResponses})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product.ToResponse()})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updateData struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Quantity    int     `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	if updateData.Price != 0 {
		product.Price = updateData.Price
	}
	if updateData.Quantity != 0 {
		product.Quantity = updateData.Quantity
	}

	if err := pc.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product.ToResponse()})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.ImagePath != "" {
		utils.DeleteFile(product.ImagePath)
	}

	if err := pc.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (pc *ProductController) UploadProductImage(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	if err := utils.ValidateFileType(file, allowedTypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product.ImagePath != "" {
		utils.DeleteFile(product.ImagePath)
	}

	imagePath, err := utils.UploadFile(c, file, "products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	product.ImagePath = imagePath
	if err := pc.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product.ToResponse()})
} 
