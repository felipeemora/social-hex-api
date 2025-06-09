package users

import (
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/domain"
	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
	common "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
	exceptions "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/exceptions"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

type ActivateUserHandler struct {
	logger              logger.Logger
	activateUserUsecase in.ActivateUserPort
}

func NewActivateUserHandler(activateUserUsecase in.ActivateUserPort, logger logger.Logger) *ActivateUserHandler {
	return &ActivateUserHandler{
		logger:              logger,
		activateUserUsecase: activateUserUsecase,
	}
}

func (h *ActivateUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	invitationID := chi.URLParam(r, "invitationID")

	err := h.activateUserUsecase.Execute(r.Context(), &invitationID)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			exceptions.BadRequestError(w, r, h.logger, err, nil)
		default:
			exceptions.InternalServerError(w, r, h.logger, err, nil)
		}
		return
	}

	h.logger.Info("User activated successfully")
	if err := common.JsonResponse(w, http.StatusOK, "User Activated"); err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}
}

// ActivateUserHandler godoc
//
//	@Summary		Activate a registered user
//	@Description	Activate a registered user using an invitation ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			invitationID	path		string	true	"Invitation ID"
//	@Success		200				{string}	string	"User activated"
//	@Failure		400				{object}	common.ErrorAPIResponse
//	@Failure		500				{object}	common.ErrorAPIResponse
//	@Router			/users/activate/{invitationID} [put]
func (h *ActivateUserHandler) RegisterRoutes(router chi.Router) {
	router.Put("/users/activate/{invitationID}", h.Handler)
}
