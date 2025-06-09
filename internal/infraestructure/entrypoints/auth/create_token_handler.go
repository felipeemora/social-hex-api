package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/domain"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/auth/dto"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/auth/mapper"
	common "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
	exceptions "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/exceptions"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

type CreateTokenHandler struct {
	logger             logger.Logger
	createTokenUsecase in.CreateTokenPort
}

func NewCreateTokenHandler(logger logger.Logger, createTokenUsecase in.CreateTokenPort) *CreateTokenHandler {
	return &CreateTokenHandler{
		logger:             logger,
		createTokenUsecase: createTokenUsecase,
	}
}

func (h *CreateTokenHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreateTokenRequest
	if err := common.ReadJson(w, r, &payload); err != nil {
		exceptions.BadRequestError(w, r, h.logger, err, nil)
		return
	}

	if errs := common.ValidateStruct(&payload); len(errs) > 0 {
		exceptions.BadRequestError(w, r, h.logger, errors.New("error in fields validation"), errs)
		return
	}

	userDomain := mapper.ToDomain(&payload)
	err := h.createTokenUsecase.Execute(r.Context(), userDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			exceptions.UnauthorizedError(w, r, h.logger, fmt.Errorf("invalid credentials"), nil)
		default:
			exceptions.InternalServerError(w, r, h.logger, err, nil)
		}
		return
	}

	if err := common.JsonResponse(w, http.StatusCreated, mapper.FromDomain(userDomain)); err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}
}

// Create Token Handler
//
//	@Summary		Create Token
//	@Description	create a new token for a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		dto.CreateTokenRequest	true	"User token creation payload"
//	@Success		201		{object}	dto.CreateTokenResponse	"User token created successfully"
//	@Failure		400		{object}	common.ErrorAPIResponse
//	@Failure		401		{object}	common.ErrorAPIResponse
//	@Failure		500		{object}	common.ErrorAPIResponse
//	@Router			/authentication/token [post]
func (h *CreateTokenHandler) RegisterRoutes(router chi.Router) {
	router.Post("/authentication/token", h.Handler)
}
