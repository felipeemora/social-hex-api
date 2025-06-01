package internal

import (
	"database/sql"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/application"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints"
	createpost "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/swagger"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

type Handlers interface {
	Handler(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router chi.Router)
}

func BuildDependencies(db *sql.DB, logger logger.Logger) *[]Handlers {
	logger.Info("Building dependencies...")
	drivenadaptersDeps := drivenadapters.NewDrivenAdaptersDependencies(db)
	applicationDeps := application.NewApplicationDependencies(drivenadaptersDeps)
	entrypointsDeps := entrypoints.NewEntrypointsDependencies(applicationDeps)
	logger.Info("Dependencies built successfully")

	logger.Info("Creating handlers...")
	swaggerHander := swagger.NewSwaggerHandler()

	postHandler := createpost.NewCreatePostHandler(entrypointsDeps.CreatePostUsecase, logger)
	logger.Info("Handlers created successfully")

	return &[]Handlers{
		swaggerHander,
		postHandler,
	}
}
