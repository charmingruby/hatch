package endpoint

import (
	"context"
	"time"

	"github.com/charmingruby/pack/internal/shared/http/rest"
	"github.com/gin-gonic/gin"
)

const TIMEOUT_IN_SECONDS = 10

func (e *Endpoint) readinessHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_IN_SECONDS*time.Second)
	defer cancel()

	if err := e.service.db.Ping(ctx); err != nil {
		rest.SendServiceUnavailableResponse(c, "database")
		return
	}

	rest.SendOKResponse(c, "", nil)
}
