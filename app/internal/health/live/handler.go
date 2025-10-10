package live

import (
	"HATCH_APP/internal/shared/transport/http"
	"HATCH_APP/pkg/logger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	log *logger.Logger
	r   *gin.Engine
}

func registerRoute(h handler) {
	api := h.r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/livez", h.handle)
}

func (h *handler) handle(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.InfoContext(ctx, "endpoint/Liveness: request received")
	h.log.InfoContext(ctx, "endpoint/Liveness: finished successfully")

	http.SendOKResponse(c, "", nil)
}
