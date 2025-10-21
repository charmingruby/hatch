package rest

import (
	"HATCH_APP/pkg/db/postgres"
	"HATCH_APP/pkg/telemetry"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

const timeoutInSeconds = 10

func registerProbes(log *telemetry.Logger, r *gin.Engine, db *postgres.Client) {
	r.GET("/livez", livenessRoute(log))
	r.GET("/readyz", readinessRoute(log, db))
}

func livenessRoute(log *telemetry.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Liveness: request received")
		log.InfoContext(ctx, "endpoint/Liveness: finished successfully")

		SendOKResponse(c, "", nil)
	}
}

func readinessRoute(log *telemetry.Logger, db *postgres.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Readiness: request received")

		ctx, cancel := context.WithTimeout(ctx, timeoutInSeconds*time.Second)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err.Error())

			SendServiceUnavailableResponse(c, "database")
			return
		}

		log.InfoContext(ctx, "endpoint/Readiness: finished successfully")

		SendOKResponse(c, "", nil)
	}
}
