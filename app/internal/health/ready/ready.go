package ready

import (
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

func New(
	log *logger.Logger,
	router *gin.Engine,
	db *postgres.Client,
) {
	registerRoute(handler{
		log: log,
		r:   router,
		db:  db,
	})
}
