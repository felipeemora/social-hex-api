package out

import "time"

type ConfigurationsPort interface {
	GetActivationURL(string) string
	GetInvitationTemplateFile() string
	GetEnvironment() string
	GetInvitationExpiration() time.Duration
	GetJWTTokenExpiration() time.Duration
}
