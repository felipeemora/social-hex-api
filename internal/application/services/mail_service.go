package services

import (
	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
)

var (
	productionEnvironment = "production"
)

type MailService struct {
	logger             logger.Logger
	MailPort           out.MailPort
	configurationsPort out.ConfigurationsPort
}

func NewMailService(mailPort out.MailPort, logger logger.Logger, configurationsPort out.ConfigurationsPort) *MailService {
	return &MailService{
		MailPort:           mailPort,
		logger:             logger,
		configurationsPort: configurationsPort,
	}
}

func (ms *MailService) Send(username, email, invitationID string) error {
	templateFile := ms.configurationsPort.GetInvitationTemplateFile()
	emailVars := ms.getEmailVars(username, invitationID)
	isSandbox := ms.configurationsPort.GetEnvironment() != productionEnvironment

	statusCode, err := ms.MailPort.Send(templateFile, username, email, emailVars, isSandbox)
	if err != nil {
		ms.logger.Error("error sending email", "error", err)
		return err
	}

	ms.logger.Infow("email sent", "status code", statusCode, "email", email)
	return nil
}

func (ms *MailService) getEmailVars(username, invitationID string) any {
	return struct {
		Username      string
		ActivationURL string
	}{
		Username:      username,
		ActivationURL: ms.configurationsPort.GetActivationURL(invitationID),
	}
}
