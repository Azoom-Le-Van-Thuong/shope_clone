package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
