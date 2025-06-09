package usecases

import (
	"context"
	"time"

	"github.com/felipeemora/social-hex-api/internal/application/services"
	domain "github.com/felipeemora/social-hex-api/internal/domain"
	domainModels "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type CreateTokenUsecase struct {
	logger             logger.Logger
	tokenService       TokenService
	storagePort        out.StorageUserPort
	configurationsPort out.ConfigurationsPort
}

func NewCreateTokenUsecase(logger logger.Logger, storagePort out.StorageUserPort, configurationsPort out.ConfigurationsPort, tokenService TokenService) *CreateTokenUsecase {
	return &CreateTokenUsecase{
		logger:             logger,
		tokenService:       tokenService,
		storagePort:        storagePort,
		configurationsPort: configurationsPort,
	}
}

func (uc *CreateTokenUsecase) Execute(ctx context.Context, user *domainModels.UserModel) error {
	userDB, err := uc.storagePort.GetByEmail(ctx, &user.Email)
	if err != nil {
		return err
	}

	uc.logger.Info("Validating user credentials")
	if ok := services.Compare(userDB.PasswordHash, user.Password); !ok {
		return domain.ErrInvalidCredentials
	}

	tokenExpiration := uc.configurationsPort.GetJWTTokenExpiration()

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(tokenExpiration).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	}

	token, err := uc.tokenService.GenerateToken(claims)
	if err != nil {
		return err
	}

	user.Token = &token
	uc.logger.Infow("Token generated successfully", "email", user.Email)

	return nil
}
