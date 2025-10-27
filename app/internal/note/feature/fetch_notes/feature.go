package fetch_notes

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
)

func NewFeature(
	log *telemetry.Logger,
	router *gin.RouterGroup,
	repo domain.NoteRepository,
) {
	usecase := NewService(repo)

	handler := NewHTTPHandler(log, usecase)

	router.GET("", handler)
}
