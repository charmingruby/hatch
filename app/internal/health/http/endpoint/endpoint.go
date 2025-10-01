package endpoint

import (
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	log     *logger.Logger
	router  *gin.Engine
	service service
}

type service struct {
	db *postgres.Client
}

func New(
	log *logger.Logger,
	r *gin.Engine,
	db *postgres.Client,
) *Endpoint {
	return &Endpoint{
		log:    log,
		router: r,
		service: service{
			db: db,
		},
	}
}

func (e *Endpoint) Register() {
	api := e.router.Group("/api")

	v1 := api.Group("/v1")

	v1.GET("/health/live", e.Liveness)
	v1.GET("/health/ready", e.Readiness)
}
