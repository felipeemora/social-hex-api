package dto

type CreateTokenRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}
