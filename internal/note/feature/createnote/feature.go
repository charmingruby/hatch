package createnote

import (
	"HATCH_APP/internal/note/domain"
	"net/http"
)

func Route(repo domain.NoteRepository) http.HandlerFunc {
	service := NewService(repo)

	handler := NewHandler(service)

	return handler
}
