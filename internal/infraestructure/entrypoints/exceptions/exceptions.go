package entrypoints

import (
	"net/http"

	helpers "github.com/felipeemora/social-hex-api/internal/infraestructure/entrypoints/common"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Errorw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusNotFound, "resource not found")
}

func ConflictError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("conflict error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusConflict, err.Error())
}

func UnauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusUnauthorized, err.Error())
}

func UnauthorizedBasicError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)

	helpers.WriteJSONError(w, http.StatusUnauthorized, err.Error())
}

func ForbiddenError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("forbidden error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	helpers.WriteJSONError(w, http.StatusForbidden, err.Error())
}
