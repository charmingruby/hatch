package rest

import (
	"HATCH_APP/pkg/telemetry/logger"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func registerProbes(r *gin.Engine, db *sqlx.DB) {
	r.GET("/livez", livenessRoute())
	r.GET("/readyz", readinessRoute(db))
}

func livenessRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.FromContext(c)

		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Liveness: request received")
		log.InfoContext(ctx, "endpoint/Liveness: finished successfully")

		SendOKResponse(c, "", nil)
	}
}

func readinessRoute(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.FromContext(c)

		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/Readiness: request received")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err)

			SendServiceUnavailableResponse(c, "database")
			return
		}

		log.InfoContext(ctx, "endpoint/Readiness: finished successfully")

		SendOKResponse(c, "", nil)
	}
}
