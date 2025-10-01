package endpoint

import (
	"HATCH_APP/internal/note/usecase"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	log     *logger.Logger
	router  *gin.Engine
	val     *validator.Validator
	service usecase.Service
}

func New(
	log *logger.Logger,
	router *gin.Engine,
	validator *validator.Validator,
	service usecase.Service,
) *Endpoint {
	return &Endpoint{
		log:     log,
		router:  router,
		val:     validator,
		service: service,
	}
}

func (e *Endpoint) Register() {
	api := e.router.Group("/api")

	v1 := api.Group("/v1")

	notes := v1.Group("/notes")
	notes.POST("/", e.CreateNote)
	notes.GET("/", e.ListNotes)
	notes.PATCH("/:id", e.ArchiveNote)
}
