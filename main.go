package main

import (
	"log"
	"my-app/database"
	"my-app/middleware"
	"my-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "my-app/docs"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"my-app/migration"
	"my-app/utils"
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
	migration.RunAll()

	// Setup Gin
	r := gin.Default()

	// ✅ Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React app
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.SetTrustedProxies(nil)

	r.Use(middleware.Logger())

	utils.InitFirebase()

	// Register routes
	routes.AuthRoutes(r)
	routes.EmployeeRoutes(r)
	routes.DepartmentRoutes(r)
	routes.EmploymentTypeRoutes(r)

	// Swagger route
	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Println("📘 Swagger enabled at /swagger/index.html")
	} else {
		log.Println("🚫 Swagger disabled in production")
	}

	// Run server
	log.Println("🚀 Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
