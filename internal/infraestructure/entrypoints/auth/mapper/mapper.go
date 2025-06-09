package mapper

import (
	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/auth/dto"
)

func ToDomain(dto *dto.CreateTokenRequest) *domain.UserModel {
	return &domain.UserModel{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func FromDomain(user *domain.UserModel) *dto.CreateTokenResponse {
	return &dto.CreateTokenResponse{
		Token: *user.Token,
	}
}
