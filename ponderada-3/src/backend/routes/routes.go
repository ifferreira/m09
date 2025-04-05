package routes

import (
	"github.com/gin-gonic/gin"

	"backend/controllers"
	"backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()
	productController := controllers.NewProductController()

	r.POST("/auth/register", userController.Register)
	r.POST("/auth/login", userController.Login)

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/me", userController.GetProfile)
		userRoutes.PUT("/me", userController.UpdateProfile)
		userRoutes.POST("/me/image", userController.UploadProfileImage)
	}

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware())
	{
		adminRoutes.GET("/users", userController.GetAllUsers)
		adminRoutes.GET("/users/:id", userController.GetUserByID)
		adminRoutes.DELETE("/users/:id", userController.DeleteUser)
	}

	r.GET("/products", productController.GetAllProducts)
	r.GET("/products/:id", productController.GetProductByID)

	protectedProducts := r.Group("/products")
	protectedProducts.Use(middleware.AuthMiddleware())
	{
		protectedProducts.POST("", productController.CreateProduct)
		protectedProducts.PUT("/:id", productController.UpdateProduct)
		protectedProducts.DELETE("/:id", productController.DeleteProduct)
		protectedProducts.POST("/:id/image", productController.UploadProductImage)
	}
}

func InitRoutes(r *gin.Engine) {
	SetupRoutes(r)
} 
