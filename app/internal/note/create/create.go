package create

import (
	"HATCH_APP/internal/note/shared/repository"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

func New(
	log *logger.Logger,
	router *gin.Engine,
	validator *validator.Validator,
	repo repository.NoteRepo,
) {
	uc := NewUseCase(repo)

	registerRoute(handler{
		log: log,
		r:   router,
		val: validator,
		svc: uc,
	})
}
