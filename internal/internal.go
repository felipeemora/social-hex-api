package internal

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/application"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints"
	createpost "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts"
	"github.com/go-chi/chi/v5"
)

type Handlers interface {
	Handler(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *chi.Mux)
}

func BuildDependencies(db *sql.DB) *[]Handlers {
	log.Println("Building dependencies...")
	drivenadaptersDeps := drivenadapters.NewDrivenAdaptersDependencies(db)
	applicationDeps := application.NewApplicationDependencies(drivenadaptersDeps)
	entrypointsDeps := entrypoints.NewEntrypointsDependencies(applicationDeps)
	log.Println("Dependencies built successfully")

	log.Print("Creating handlers...")
	postHandler := createpost.NewCreatePostHandler(entrypointsDeps.CreatePostUsecase)
	log.Println("Handlers created successfully")

	return &[]Handlers{
		postHandler,
	}
}
