package mapper

import (
	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts/dto"
)

func ToDomain(dto *dto.CreatePostRequestDTO, userID int) *domain.PostModel {
	return &domain.PostModel{
		Title:   dto.Title,
		Content: dto.Content,
		Tags:    dto.Tags,
		UserID:  userID,
	}
}

func FromDomain(postModel *domain.PostModel) *dto.PostResponseDTO {
	return &dto.PostResponseDTO{
		ID:      postModel.ID,
		CreatePostRequestDTO: &dto.CreatePostRequestDTO{
			Title:   postModel.Title,
			Content: postModel.Content,
			Tags:    postModel.Tags,
			UserID:  postModel.UserID,
		},
	}
}
