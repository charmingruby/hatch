package endpoint

import (
	"HATCH_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) Liveness(c *gin.Context) {
	ctx := c.Request.Context()
	e.log.InfoContext(ctx, "endpoint/Liveness: request received")
	e.log.InfoContext(ctx, "endpoint/Liveness: finished successfully")

	rest.SendOKResponse(c, "", nil)
}
