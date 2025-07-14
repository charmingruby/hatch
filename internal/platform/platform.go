package platform

import (
	"github/charmingruby/pack/internal/platform/delivery/rest"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	rest.RegisterRoutes(r)
}
