package listnotes

import (
	"HATCH_APP/internal/note/domain"
	"net/http"
)

func New(repo domain.NoteRepository) http.HandlerFunc {
	service := NewService(repo)

	handler := NewHandler(service)

	return handler
}
