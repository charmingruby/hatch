package endpoint

import (
	"HATCH_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) livenessHandler(c *gin.Context) {
	rest.SendOKResponse(c, "", nil)
}
