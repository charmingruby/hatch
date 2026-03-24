package httpx

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteOKResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusOK, v)
}

func WriteCreatedResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusCreated, v)
}

func WriteBadRequestResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusBadRequest, v)
}

func WriteNotFoundResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusNotFound, v)
}

func WriteServiceUnavailableResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusServiceUnavailable, v)
}

func WriteConflictResponse(w http.ResponseWriter, v any) {
	writeResponse(w, http.StatusConflict, v)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusInternalServerError, ErrorResponse{
		Message: "Internal Server Error",
	})
}

func WriteEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
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
