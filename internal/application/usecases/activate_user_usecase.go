package usecases

import (
	"context"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
)

type ActivateUserUsecase struct {
	storagePort out.StorageUserPort
}

func NewActivateUserUsecase(storagePort out.StorageUserPort) *ActivateUserUsecase {
	return &ActivateUserUsecase{
		storagePort: storagePort,
	}
}

func (uc *ActivateUserUsecase) Execute(ctx context.Context, invitationID *string) error {
	if err := uc.storagePort.Activate(ctx, invitationID); err != nil {
		return err
	}

	return nil
}
