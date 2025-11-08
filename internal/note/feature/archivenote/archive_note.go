package archivenote

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
)

func NewFeature(log *telemetry.Logger, repo domain.NoteRepository) gin.HandlerFunc {
	usecase := NewService(repo)

	handler := NewHTTPHandler(log, usecase)

	return handler
}
