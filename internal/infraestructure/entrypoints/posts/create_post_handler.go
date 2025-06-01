package createpost

import (
	"errors"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
	common "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
	exceptions "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/exceptions"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts/dto"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/posts/mapper"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

type CreatePostHandler struct {
	logger            logger.Logger
	CreatePostUsecase in.CreatePostPort
}

func NewCreatePostHandler(createPostUsecase in.CreatePostPort, logger logger.Logger) *CreatePostHandler {
	return &CreatePostHandler{
		logger:            logger,
		CreatePostUsecase: createPostUsecase,
	}
}

func (h *CreatePostHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreatePostRequestDTO
	if err := common.ReadJson(w, r, &payload); err != nil {
		exceptions.BadRequestError(w, r, h.logger, err, nil)
		return
	}

	if errs := common.ValidateStruct(&payload); len(errs) > 0 {
		exceptions.BadRequestError(w, r, h.logger, errors.New("error in fields validation"), errs)
		return
	}

	postDomain := mapper.ToDomain(&payload, 1)
	ctx := r.Context()

	postDomain, err := h.CreatePostUsecase.Execute(ctx, postDomain)
	h.logger.Infow("Post created successfully", "post", postDomain)

	if err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}

	if err := common.JsonResponse(w, http.StatusCreated, mapper.FromDomain(postDomain)); err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}
}

// ListAccounts lists all existing accounts
//
//	@Summary		Create Post
//	@Description	create a new post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		dto.CreatePostRequestDTO	true	"Post creation payload"
//	@Success		201		{object}	dto.PostSuccessAPIResponse "Post created successfully"
//	@Failure		400		{object}	common.ErrorAPIResponse
//	@Failure		500		{object}	common.ErrorAPIResponse
//	@Router			/posts [post]
func (h *CreatePostHandler) RegisterRoutes(router chi.Router) {
	router.Post("/posts", h.Handler)
}
