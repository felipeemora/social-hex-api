package mailtrap

import (
	"embed"

	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
)

const (
	maxRetries = 3
	FromName   = "SocialAPI"
)

//go:embed templates/*
var FS embed.FS

type MailTrapClient struct {
	fromEmail string
	apiKey    string
	host      string
	port      int
	username  string
}

func NewMailTrapClient(mailConfig *configurations.MailConfig) *MailTrapClient {
	return &MailTrapClient{
		fromEmail: mailConfig.FromEmail,
		apiKey:    mailConfig.ApiKey,
		host:      mailConfig.Host,
		port:      mailConfig.Port,
		username:  mailConfig.Username,
	}
}
