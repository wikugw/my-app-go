package handlers

type GoogleLoginRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}
