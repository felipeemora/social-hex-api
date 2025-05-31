package application

import (
	"github.com/felipeemora/social-hex-api/internal/application/usecases"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters"
)

type ApplicationDependencies struct {
	CreatePostUsecase *usecases.CreatePostUsecase
}

func NewApplicationDependencies(drivenadaptersDeps *drivenadapters.DrivenAdaptersDependencies) *ApplicationDependencies {
	return &ApplicationDependencies{
		CreatePostUsecase: usecases.NewCreatePostUseCase(drivenadaptersDeps.StoragePort),
	}
}
