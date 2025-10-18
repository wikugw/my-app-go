package handlers

type GoogleLoginRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}
