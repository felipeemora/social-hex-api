package entrypoints

import (
	"github.com/felipeemora/social-hex-api/internal/application"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
)

type EntrypointsDependencies struct {
	CreatePostUsecase in.CreatePostPort
	CreateUserUsecase in.CreateUserPort
}

func NewEntrypointsDependencies(applicationDeps *application.ApplicationDependencies) *EntrypointsDependencies {
	return &EntrypointsDependencies{
		CreatePostUsecase: applicationDeps.CreatePostUsecase,
		CreateUserUsecase: applicationDeps.CreateUserUsecase,
	}
}
