package out

import (
	"context"
	"time"

	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
)

type StorageUserPort interface {
	CreateUserAndInvite(context.Context, *domain.UserModel, string, time.Duration) error
	Delete(context.Context, int64) error
}
