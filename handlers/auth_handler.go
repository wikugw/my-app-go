package handlers

import (
	"context"
	"net/http"
	"os"

	"my-app/types/handlers"
	"my-app/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func GoogleLoginHandler(c *gin.Context) {
	var req handlers.GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id_token"})
		return
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")

	// verifikasi token dari Google
	payload, err := idtoken.Validate(context.Background(), req.IDToken, clientID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	token, err := utils.GenerateJWT(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"user": gin.H{
			"email": email,
			"name":  name,
		},
	})
}
