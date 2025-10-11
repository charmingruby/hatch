package health

import (
	"HATCH_APP/internal/shared/transport/http"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

const timeoutInS = 10

func registerRoutes(
	log *logger.Logger,
	r *gin.Engine,
	db *postgres.Client,
) {
	r.GET("/livez", livenessRoute(log))
	r.GET("/readyz", readinessRoute(log, db))
}

func livenessRoute(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Liveness: request received")
		log.InfoContext(ctx, "endpoint/Liveness: finished successfully")

		http.SendOKResponse(c, "", nil)
	}
}

func readinessRoute(log *logger.Logger, db *postgres.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Readiness: request received")

		ctx, cancel := context.WithTimeout(ctx, timeoutInS*time.Second)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err.Error())

			http.SendServiceUnavailableResponse(c, "database")
			return
		}

		log.InfoContext(ctx, "endpoint/Readiness: finished successfully")

		http.SendOKResponse(c, "", nil)
	}
}
