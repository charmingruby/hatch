package archive

import (
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

func New(
	log *logger.Logger,
	api *gin.RouterGroup,
	repo repository.NoteRepo,
) {
	registerRoute(
		log,
		api,
		NewUseCase(repo),
	)
}
