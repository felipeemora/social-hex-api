package out

import (
	"context"

	"github.com/felipeemora/social-hex-api/internal/domain/models"
)

type StoragePostPort interface {
	Save(context.Context, *domain.PostModel) (*domain.PostModel, error)
}
