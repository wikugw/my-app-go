package routes

import (
	"my-app/handlers"
	"my-app/middleware"

	"github.com/gin-gonic/gin"
)

func RecruitmentRoutes(r *gin.Engine) {
	// Middleware khusus department routes
	deptGroup := r.Group("/recruitments")
	deptGroup.Use(middleware.Logger())
	{
		deptGroup.POST("", handlers.CreateRecruitmentHandler)
	}
}
