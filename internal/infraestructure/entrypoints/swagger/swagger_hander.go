package swagger

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerHandler struct {
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (h *SwaggerHandler) Handler(w http.ResponseWriter, r *http.Request) {
	httpSwagger.Handler(httpSwagger.URL(":8080/swagger/doc.json")).ServeHTTP(w, r)
}

func (h *SwaggerHandler) RegisterRoutes(router chi.Router) {
	router.Get("/swagger/*", h.Handler)
}
