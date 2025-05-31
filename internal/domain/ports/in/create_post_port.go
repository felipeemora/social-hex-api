package in

import (
	"context"

	"github.com/felipeemora/social-hex-api/internal/domain/models"
)

type CreatePostPort interface {
	Execute(context.Context, *domain.PostModel) (*domain.PostModel, error)
}
