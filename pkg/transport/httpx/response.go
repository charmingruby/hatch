package httpx

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteOKResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusOK, r)
}

func WriteCreatedResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusCreated, r)
}

func WriteBadRequestResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusBadRequest, r)
}

func WriteNotFoundResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusNotFound, r)
}

func WriteServiceUnavailableResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusServiceUnavailable, r)
}

func WriteConflictResponse(w http.ResponseWriter, r Response) {
	writeResponse(w, http.StatusInternalServerError, r)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusInternalServerError, Response{
		Message: "Internal Server Error",
		Data:    nil,
	})
}

func WriteEmptyResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusNoContent, Response{})
}

func writeResponse(w http.ResponseWriter, status int, r Response) {
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
