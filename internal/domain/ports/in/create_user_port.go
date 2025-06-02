package in

import (
	"context"

	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
)

type CreateUserPort interface {
	Execute(context.Context, *domain.UserModel) error
}