package routes

import (
	"my-app/handlers"
	"my-app/middleware"

	"github.com/gin-gonic/gin"
)

func RecruitmentRoutes(r *gin.Engine) {
	// Middleware khusus department routes
	deptGroup := r.Group("/recruitments")
	deptGroup.Use(middleware.Logger()).Use(middleware.AuthMiddleware())
	{
		deptGroup.POST("", handlers.CreateRecruitmentHandler)
		deptGroup.GET("/active", handlers.GetActiveRecruitmentsHandler)
	}
}
