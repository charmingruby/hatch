package archivenote

import (
	"HATCH_APP/internal/note/domain"

	"github.com/gin-gonic/gin"
)

func NewFeature(repo domain.NoteRepository) gin.HandlerFunc {
	usecase := NewService(repo)

	handler := NewHTTPHandler(usecase)

	return handler
}
