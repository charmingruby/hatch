package endpoint

import (
	"PACK_APP/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	router  *gin.Engine
	service service
}

type service struct {
	db *postgres.Client
}

func New(r *gin.Engine, db *postgres.Client) *Endpoint {
	return &Endpoint{
		router: r,
		service: service{
			db: db,
		},
	}
}

func (e *Endpoint) Register() {
	api := e.router.Group("/api")

	v1 := api.Group("/v1")

	v1.GET("/health/live", e.livenessHandler)
	v1.GET("/health/ready", e.readinessHandler)
}
