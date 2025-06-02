package drivenadapters

import (
	"database/sql"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/mailtrap"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/postgres"
)

type DrivenAdaptersDependencies struct {
	StoragePostPort    out.StoragePostPort
	StorageUserPort    out.StorageUserPort
	ConfigurationsPort out.ConfigurationsPort
	MailPort           out.MailPort
}

func NewDrivenAdaptersDependencies(db *sql.DB, mailConfig *configurations.MailConfig) *DrivenAdaptersDependencies {
	return &DrivenAdaptersDependencies{
		StoragePostPort:    postgres.NewPostgresAdapter(db),
		StorageUserPort:    postgres.NewStorageUserAdapter(db),
		ConfigurationsPort: configurations.NewConfigurationsAdapter(),
		MailPort:           mailtrap.NewMailTrapClient(mailConfig),
	}
}
