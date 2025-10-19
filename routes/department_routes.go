package routes

import (
	"my-app/handlers"
	"my-app/middleware"

	"github.com/gin-gonic/gin"
)

func DepartmentRoutes(r *gin.Engine) {
	// Middleware khusus department routes
	deptGroup := r.Group("/departments")
	deptGroup.Use(middleware.Logger()).Use(middleware.AuthMiddleware())
	{
		deptGroup.GET("", handlers.GetDepartmentsHandler)
	}
}
