package endpoint

import (
	"context"
	"time"

	"HATCH_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

const TIMEOUT_IN_SECONDS = 10

func (e *Endpoint) Readiness(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/Readiness: request received")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT_IN_SECONDS*time.Second)
	defer cancel()

	if err := e.service.db.Ping(ctx); err != nil {
		e.log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err.Error())

		rest.SendServiceUnavailableResponse(c, "database")
		return
	}

	e.log.InfoContext(ctx, "endpoint/Readiness: finished successfully")

	rest.SendOKResponse(c, "", nil)
}
