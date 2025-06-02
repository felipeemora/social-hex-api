package application

import (
	"github.com/felipeemora/social-hex-api/internal/application/services"
	"github.com/felipeemora/social-hex-api/internal/application/usecases"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
)

type ApplicationDependencies struct {
	CreatePostUsecase *usecases.CreatePostUsecase
	CreateUserUsecase *usecases.CreateUserUsecase
}

func NewApplicationDependencies(drivenadaptersDeps *drivenadapters.DrivenAdaptersDependencies, logger logger.Logger) *ApplicationDependencies {
	return &ApplicationDependencies{
		CreatePostUsecase: usecases.NewCreatePostUseCase(
			drivenadaptersDeps.StoragePostPort,
		),
		CreateUserUsecase: usecases.NewCreateUserUsecase(
			logger,
			drivenadaptersDeps.StorageUserPort,
			services.NewMailService(drivenadaptersDeps.MailPort, logger, drivenadaptersDeps.ConfigurationsPort),
			drivenadaptersDeps.ConfigurationsPort,
		),
	}
}
