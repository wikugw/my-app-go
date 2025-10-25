package routes

import (
	"my-app/handlers"
	"my-app/middleware"

	"github.com/gin-gonic/gin"
)

func EmploymentTypeRoutes(r *gin.Engine) {
	deptGroup := r.Group("/employment-types")
	deptGroup.Use(middleware.Logger()).Use(middleware.AuthMiddleware())
	{
		deptGroup.GET("", handlers.GetEmploymentTypesHandler)
	}
}
