package live

import (
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

func New(
	log *logger.Logger,
	router *gin.Engine,
) {
	registerRoute(handler{
		log: log,
		r:   router,
	})
}
