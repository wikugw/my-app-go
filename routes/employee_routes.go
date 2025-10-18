package routes

import (
	"my-app/handlers"
	"my-app/middleware"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	// Middleware khusus employee routes
	employeeGroup := r.Group("/employees")
	employeeGroup.Use(middleware.Logger()).Use(middleware.AuthMiddleware())
	{
		employeeGroup.POST("", handlers.CreateEmployeeHandler)
		employeeGroup.GET("", handlers.GetEmployeeByEmailHandler)
	}
}
