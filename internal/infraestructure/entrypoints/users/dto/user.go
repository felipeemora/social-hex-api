package dto

type CreateUserRequestDTO struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	RoleName string `json:"role_name" validate:"required"`
}

type UserResponseDTO struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	IsActive     bool   `json:"is_active"`
	InvitationID string `json:"invitation_id"`
}
