package usecases

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/felipeemora/social-hex-api/internal/application/services"
	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/google/uuid"
)

type MailService interface {
	Send(username, email, invitationID string) error
}

type CreateUserUsecase struct {
	logger             logger.Logger
	mailService        MailService
	storagePort        out.StorageUserPort
	configurationsPort out.ConfigurationsPort
}

func NewCreateUserUsecase(logger logger.Logger, storagePort out.StorageUserPort, mailService MailService, configurationsPort out.ConfigurationsPort) *CreateUserUsecase {
	return &CreateUserUsecase{
		logger:      logger,
		storagePort: storagePort,
		mailService: mailService,
		configurationsPort: configurationsPort,
	}
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, userModel *domain.UserModel) error {
	var err error
	userModel.PasswordHash, err = services.Hash(userModel.Password)
	if err != nil {
		return err
	}

	invitationID, invitationIDHash := generateInvitationId()
	userModel.InvitationID = invitationID
	invitationExpiration := uc.configurationsPort.GetInvitationExpiration()

	if err := uc.storagePort.CreateUserAndInvite(ctx, userModel, invitationIDHash, invitationExpiration); err != nil {
		return err
	}

	if err := uc.mailService.Send(userModel.Username, userModel.Email, invitationID); err != nil {
		if err := uc.storagePort.Delete(ctx, userModel.ID); err != nil {
			uc.logger.Errorw("failed to delete user after email send failure", "error", err)
		}
		return err
	}

	return nil
}

func generateInvitationId() (string, string) {
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	return plainToken, hex.EncodeToString(hash[:])
}
