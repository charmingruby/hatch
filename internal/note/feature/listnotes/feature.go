package listnotes

import (
	"HATCH_APP/internal/note/domain"

	"github.com/gin-gonic/gin"
)

func New(
	repo domain.NoteRepository,
) gin.HandlerFunc {
	service := NewService(repo)

	handler := NewHandler(service)

	return handler
}
