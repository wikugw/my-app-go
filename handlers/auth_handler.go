package handlers

import (
	"context"
	"net/http"

	"my-app/types/handlers"
	"my-app/utils"

	"github.com/gin-gonic/gin"
)

// GoogleLoginHandler godoc
// @Summary Login menggunakan Google OAuth
// @Description Verifikasi `id_token` dari Google Sign-In dan menghasilkan JWT internal untuk autentikasi selanjutnya.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body handlers.GoogleLoginRequest true "Google ID Token"
// @Success 200 {object} map[string]interface{} "JWT token dan data pengguna"
// @Failure 400 {object} map[string]string "Missing id_token"
// @Failure 401 {object} map[string]string "Invalid Google token"
// @Failure 500 {object} map[string]string "Failed to create token"
// @Router /auth/google [post]
func GoogleLoginHandler(c *gin.Context) {
	var req handlers.GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.IDToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id_token"})
		return
	}

	// verifikasi token Firebase
	token, err := utils.FirebaseAuthClient.VerifyIDToken(context.Background(), req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)

	jwtToken, err := utils.GenerateJWT(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": jwtToken,
		"user": gin.H{
			"email": email,
			"name":  name,
		},
	})
}
