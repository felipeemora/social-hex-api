package main

import (
	"log"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/postgres"
	"github.com/go-chi/chi/v5"
)

type application struct {
	cfg *configurations.Configurations
}

func main() {
	app := &application{
		cfg: configurations.Load(),
	}

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

	log.Fatal(app.run(mux))
}

func (app *application) mount(handlers *[]internal.Handlers) http.Handler {
	mux := chi.NewRouter()

	for _, handler := range *handlers {
		handler.RegisterRoutes(mux)
	}

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
