package endpoint

import (
	"github.com/charmingruby/pack/internal/shared/http/rest"
	"github.com/gin-gonic/gin"
)

func (e *Endpoint) livenessHandler(c *gin.Context) {
	rest.SendOKResponse(c, "", nil)
}
