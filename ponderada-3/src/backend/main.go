package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/models"
	"backend/routes"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}

	userImagesDir := filepath.Join(uploadsDir, "users")
	productImagesDir := filepath.Join(uploadsDir, "products")
	
	if err := os.MkdirAll(userImagesDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create user images directory: %v", err)
	}
	
	if err := os.MkdirAll(productImagesDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create product images directory: %v", err)
	}

	r.Static("/uploads", "./uploads")

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a test route",
		})
	})

	initDB()

	routes.InitRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDB() {
	config.ConnectDB()
	
	db := config.GetDB()
	db.AutoMigrate(&models.User{}, &models.Product{})
	
	log.Println("Database models migrated successfully")
} 
