package createpost

import (
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
	helpers "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
	exceptions "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/exceptions"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts/dto"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts/mapper"
	"github.com/go-chi/chi/v5"
)

type CreatePostHandler struct {
	CreatePostUsecase in.CreatePostPort
}

func NewCreatePostHandler(createPostUsecase in.CreatePostPort) *CreatePostHandler {
	return &CreatePostHandler{
		CreatePostUsecase: createPostUsecase,
	}
}

func (h *CreatePostHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreatePostRequestDTO
	if err := helpers.ReadJson(w, r, &payload); err != nil {
		exceptions.BadRequestError(w, r, err)
		return
	}

	if err := helpers.Validate.Struct(payload); err != nil {
		exceptions.BadRequestError(w, r, err)
		return
	}

	postDomain := mapper.ToDomain(&payload, 1)
	ctx := r.Context()

	postDomain, err := h.CreatePostUsecase.Execute(ctx, postDomain)

	if err != nil {
		exceptions.InternalServerError(w, r, err)
		return
	}

	if err := helpers.JsonResponse(w, http.StatusCreated, mapper.FromDomain(postDomain)); err != nil {
		exceptions.InternalServerError(w, r, err)
		return
	}
}

func (h *CreatePostHandler) RegisterRoutes(router *chi.Mux) {
	router.Post("/posts", h.Handler)
}
