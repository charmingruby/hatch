package httpx

import (
	"HATCH_APP/pkg/core/apperr"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Details any    `json:"details,omitempty"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func WriteOKResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusOK, v)
}

func WriteCreatedResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusCreated, v)
}

func WriteEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteError(log *slog.Logger, w http.ResponseWriter, err error) {
	if appErr, ok := errors.AsType[*apperr.Error](err); ok {
		log.Warn("application error", "type", appErr.Type, "code", appErr.Code, "message", appErr.Message)

		writeAppError(w, appErr)
		return
	}

	log.Error("internal server error", "error", err)

	writeResponse(w, http.StatusInternalServerError, ErrorResponse{
		Message: "Internal Server Error",
	})
}

func writeResponse(w http.ResponseWriter, status int, v any) {
	if v == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Internal Server Error"}`))
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func writeAppError(w http.ResponseWriter, err *apperr.Error) {
	status := mapStatus(err.Type)

	writeResponse(w, status, ErrorResponse{
		Message: err.Message,
		Code:    err.Code,
		Details: err.Details,
	})
}

func mapStatus(t apperr.ErrorType) int {
	switch t {
	case apperr.TypeNotFound:
		return http.StatusNotFound
	case apperr.TypeValidation:
		return http.StatusBadRequest
	case apperr.TypeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
