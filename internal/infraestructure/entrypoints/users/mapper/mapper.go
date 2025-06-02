package mapper

import (
	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/users/dto"
)

func ToDomain(dto *dto.CreateUserRequestDTO) *domain.UserModel {
	return &domain.UserModel{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Role: &domain.RoleModel{
			Name: dto.RoleName,
		},
	}
}

func FromDomain(userModel *domain.UserModel) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:           userModel.ID,
		Username:     userModel.Username,
		Email:        userModel.Email,
		CreatedAt:    userModel.CreatedAt,
		IsActive:     userModel.IsActive,
		InvitationID: userModel.InvitationID,
	}
}
