package rest

import (
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(log *logger.Logger, r *gin.Engine, svc service.UseCase, v *validator.Validator) {
	v1 := r.Group("/api/v1")

	v1.POST("/devices", createDeviceHandler(log, svc, v))
}
