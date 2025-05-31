package usecases

import (
	"context"

	"github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
)

type CreatePostUsecase struct {
	storagePort out.StoragePostPort
}

func NewCreatePostUseCase(storagePort out.StoragePostPort) *CreatePostUsecase {
	return &CreatePostUsecase{
		storagePort: storagePort,
	}
}

func (uc *CreatePostUsecase) Execute(ctx context.Context, postModel *domain.PostModel) (*domain.PostModel, error) {
	return uc.storagePort.Save(ctx, postModel)
}
