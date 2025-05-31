package drivenadapters

import (
	"database/sql"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/out"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/postgres"
)

type DrivenAdaptersDependencies struct {
	StoragePort out.StoragePostPort
}

func NewDrivenAdaptersDependencies(db *sql.DB) *DrivenAdaptersDependencies {
	return &DrivenAdaptersDependencies{
		StoragePort: postgres.NewPostgresAdapter(db),
	}
}
