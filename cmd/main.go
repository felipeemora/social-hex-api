package main

import (
	"log"
	"net/http"

	"github.com/felipeemora/social-hex-api/docs"
	"github.com/felipeemora/social-hex-api/internal"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/postgres"
	"github.com/go-chi/chi/v5"
)

const (
	version = "0.0.1"
 	version_path_v1 = "/v1"
)

type application struct {
	cfg *configurations.Configurations
}

//	@title			Social Hex API
//	@description	This is a server for a social media application.

//	@contact.name	Felipe Mora
//	@contact.email	felipemora@mail.com
func main() {
	app := &application{
		cfg: configurations.Load(),
	}

	app.swaggerInit()
	
	db, err := postgres.New(
		app.cfg.DBConfig.Addr,
		app.cfg.DBConfig.MaxOpenConns,
		app.cfg.DBConfig.MaxIdleConns,
		app.cfg.DBConfig.MaxIdleTime,
	)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to the database")

	handlers := internal.BuildDependencies(db)

	mux := app.mount(handlers)
	err = app.run(mux)

	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) mount(handlers *[]internal.Handlers) http.Handler {
	mux := chi.NewRouter()

	mux.Route("/v1", func(r chi.Router) {
		for _, handler := range *handlers {
			handler.RegisterRoutes(r)
		}
	})

	return mux
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:    app.cfg.ServerPort,
		Handler: mux,
	}

	log.Printf("Server listening on port %s", app.cfg.ServerPort)
	return srv.ListenAndServe()
}

func (app *application) swaggerInit() {
	docs.SwaggerInfo.BasePath = version_path_v1
	docs.SwaggerInfo.Host = app.cfg.ApiURL
	docs.SwaggerInfo.Version = version
}
