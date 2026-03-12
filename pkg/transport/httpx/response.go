package httpx

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteOKResponse(w http.ResponseWriter, msg string, data any) {
	writeResponse(w, http.StatusOK, response{
		Message: msg,
		Data:    data,
	})
}

func WriteCreatedResponse(w http.ResponseWriter, msg string, data any) {
	writeResponse(w, http.StatusCreated, response{
		Message: msg,
		Data:    data,
	})
}

func WriteBadRequestResponse(w http.ResponseWriter, msg string) {
	writeResponse(w, http.StatusBadRequest, response{
		Message: msg,
		Data:    nil,
	})
}

func WriteNotFoundResponse(w http.ResponseWriter, msg string) {
	writeResponse(w, http.StatusNotFound, response{
		Message: msg,
		Data:    nil,
	})
}

func WriteServiceUnavailableResponse(w http.ResponseWriter, msg string) {
	writeResponse(w, http.StatusServiceUnavailable, response{
		Message: msg,
		Data:    nil,
	})
}

func WriteConflictResponse(w http.ResponseWriter, msg string) {
	writeResponse(w, http.StatusInternalServerError, response{
		Message: msg,
		Data:    nil,
	})
}

func WriteInternalServerErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusInternalServerError, response{
		Message: "Internal Server Error",
		Data:    nil,
	})
}

func WriteEmptyResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusNoContent, response{})
}

func writeResponse(w http.ResponseWriter, status int, r response) {
	if r.Message == "" && r.Data == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Internal Server Error"}`))
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(body)
}
