package platform

import (
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	RegisterRoutes(r)
}
