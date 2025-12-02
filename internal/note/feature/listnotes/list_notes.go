package listnotes

import (
	"HATCH_APP/internal/note/domain"

	"github.com/gin-gonic/gin"
)

func NewFeature(
	repo domain.NoteRepository,
) gin.HandlerFunc {
	service := NewService(repo)

	handler := NewHTTPHandler(service)

	return handler
}
