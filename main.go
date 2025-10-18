package main

import (
	"log"
	"my-app/database"
	"my-app/middleware"
	"my-app/models"
	"my-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "my-app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Employee API
// @version 1.0
// @description API documentation for Employee CRUD
// @host localhost:8080
// @BasePath /
func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect DB
	database.Connect()

	// Auto-migrate models
	if err := database.DB.AutoMigrate(&models.Employee{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("🚀 Migration successful!")

	// Setup Gin
	r := gin.Default()

	r.Use(middleware.Logger())

	// Register routes
	routes.EmployeeRoutes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	log.Println("🚀 Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
