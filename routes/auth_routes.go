package routes

import (
	"my-app/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/google", handlers.GoogleLoginHandler)
}
