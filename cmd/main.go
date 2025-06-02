package main

import (
	"log"
	"net/http"

	"github.com/felipeemora/social-hex-api/docs"
	"github.com/felipeemora/social-hex-api/internal"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/postgres"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

const (
	version         = "0.0.1"
	version_path_v1 = "/v1"
)

type application struct {
	cfg    *configurations.Configurations
	logger logger.Logger
}

//	@title			Social Hex API
//	@description	This is a server for a social media application.

// @contact.name	Felipe Mora
// @contact.email	felipemora@mail.com
func main() {
	app := &application{
		cfg:    configurations.Load(),
		logger: logger.NewZapLogger(),
	}

	defer app.logger.(*logger.ZapLogger).Logger.Sync() // Sync para zap

	app.logger.Info("Starting application...")
	app.swaggerInit()

	db, err := postgres.New(
		app.cfg.DBConfig.Addr,
		app.cfg.DBConfig.MaxOpenConns,
		app.cfg.DBConfig.MaxIdleConns,
		app.cfg.DBConfig.MaxIdleTime,
	)

	if err != nil {
		app.logger.Error("Error connecting to the database: ", err)
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	app.logger.Info("Connected to the database")

	handlers := internal.BuildDependencies(db, app.logger, app.cfg.MailConfig)

	mux := app.mount(handlers)
	err = app.run(mux)

	if err != nil {
		app.logger.Error("Error running the server: ", err)
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

	app.logger.Info("Server listening on port ", app.cfg.ServerPort)
	return srv.ListenAndServe()
}

func (app *application) swaggerInit() {
	docs.SwaggerInfo.BasePath = version_path_v1
	docs.SwaggerInfo.Host = app.cfg.ApiURL
	docs.SwaggerInfo.Version = version
}
