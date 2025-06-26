package rest

import (
	"github/charmingruby/gew/internal/device/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc service.UseCase) {
	v1 := r.Group("/api/v1")

	v1.POST("/devices")
}
