package platform

import (
	"github/charmingruby/pack/internal/platform/http/rest"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	rest.New(r)
}
