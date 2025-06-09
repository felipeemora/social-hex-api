package in

import "context"

type ActivateUserPort interface {
	Execute(ctx context.Context, invitationID *string) error
}
