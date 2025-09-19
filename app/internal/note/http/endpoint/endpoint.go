package endpoint

import (
	"github.com/charmingruby/pack/internal/note/usecase"
	"github.com/charmingruby/pack/pkg/logger"
	"github.com/charmingruby/pack/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	router  *gin.Engine
	log     *logger.Logger
	val     *validator.Validator
	service usecase.Service
}

func New(
	router *gin.Engine,
	log *logger.Logger,
	validator *validator.Validator,
	service usecase.Service,
) *Endpoint {
	return &Endpoint{
		router:  router,
		log:     log,
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
