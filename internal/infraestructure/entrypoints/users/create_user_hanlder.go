package users

import (
	"errors"
	"net/http"

	"github.com/felipeemora/social-hex-api/internal/domain/ports/in"
	common "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
	exceptions "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/exceptions"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/users/dto"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/users/mapper"
	"github.com/felipeemora/social-hex-api/internal/infraestructure/support/logger"
	"github.com/go-chi/chi/v5"
)

type CreateUserHandler struct {
	logger            logger.Logger
	createUserUsecase in.CreateUserPort
}

func NewCreateUserHandler(createUserUsecase in.CreateUserPort, logger logger.Logger, ) *CreateUserHandler {
	return &CreateUserHandler{
		logger:            logger,
		createUserUsecase: createUserUsecase,
	}
}
func (h *CreateUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreateUserRequestDTO
	if err := common.ReadJson(w, r, &payload); err != nil {
		exceptions.BadRequestError(w, r, h.logger, err, nil)
		return
	}

	if errs := common.ValidateStruct(&payload); len(errs) > 0 {
		exceptions.BadRequestError(w, r, h.logger, errors.New("error in fields validation"), errs)
		return
	}

	userDomain := mapper.ToDomain(&payload)
	if err := h.createUserUsecase.Execute(r.Context(), userDomain); err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}

	h.logger.Info("User created successfully")
	if err := common.JsonResponse(w, http.StatusCreated, mapper.FromDomain(userDomain)); err != nil {
		exceptions.InternalServerError(w, r, h.logger, err, nil)
		return
	}
}

// Create User Handler
//
//	@Summary		Create User
//	@Description	create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		dto.CreateUserRequestDTO	true	"User creation payload"
//	@Success		201		{object}	dto.UserResponseDTO			"User created successfully"
//	@Failure		400		{object}	common.ErrorAPIResponse
//	@Failure		500		{object}	common.ErrorAPIResponse
//	@Router			/users [post]
func (h *CreateUserHandler) RegisterRoutes(router chi.Router) {
	router.Post("/users", h.Handler)
}
