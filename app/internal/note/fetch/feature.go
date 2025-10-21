package fetch

import (
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/pkg/telemetry"

	"github.com/gin-gonic/gin"
)

func New(log *telemetry.Logger, api *gin.RouterGroup, repo repository.NoteRepo) {
	usecase := NewUseCase(repo)
	RegisterRoute(log, api, usecase)
}
