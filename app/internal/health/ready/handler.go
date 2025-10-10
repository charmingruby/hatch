package ready

import (
	"context"
	"time"

	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/database/postgres"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

const timeout = 10

type handler struct {
	log *logger.Logger
	r   *gin.Engine
	db  *postgres.Client
}

func registerRoute(h handler) {
	api := h.r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/livez", h.handle)
}

func (h *handler) handle(c *gin.Context) {
	ctx := c.Request.Context()

	h.log.InfoContext(ctx, "endpoint/Readiness: request received")

	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		h.log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err.Error())

		rest.SendServiceUnavailableResponse(c, "database")
		return
	}

	h.log.InfoContext(ctx, "endpoint/Readiness: finished successfully")

	rest.SendOKResponse(c, "", nil)
}
