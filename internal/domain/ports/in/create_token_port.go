package in

import (
	"context"

	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
)

type CreateTokenPort interface {
	Execute(ctx context.Context, user *domain.UserModel) (error)
}
