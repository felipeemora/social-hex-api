package configurations

import (
	"fmt"
	"time"
)

type ConfigurationsAdapter struct {}

func NewConfigurationsAdapter() *ConfigurationsAdapter {
	return &ConfigurationsAdapter{}
}

func (cfg *ConfigurationsAdapter) GetActivationURL(invitationID string) string {
	return fmt.Sprintf("%s/confirm/%s", GetString("FRONTEND_URL", ""), invitationID)
}

func (cfg *ConfigurationsAdapter) GetInvitationTemplateFile() string {
	return GetString("INVITATION_TEMPLATE_FILE", "")
}

func (cfg *ConfigurationsAdapter) GetEnvironment() string {
	return GetString("ENVIRONMENT", "")
}

func (cfg *ConfigurationsAdapter) GetInvitationExpiration() time.Duration {
	expiration := GetInt("INVITATION_EXPIRATION", 0)
	if expiration <= 0 {
		return 0
	}
	return time.Hour * time.Duration(expiration)
}

func (cfg *ConfigurationsAdapter) GetJWTTokenExpiration() time.Duration{
	expiration := GetInt("JWT_TOKEN_EXPIRATION", 0)
	if expiration <= 0 {
		return 0
	}
	return time.Hour * time.Duration(expiration)
}
