package rest

import (
	"github/charmingruby/pack/internal/example/service"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(log *logger.Logger, r *gin.Engine, svc service.UseCase, v *validator.Validator) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	v1.POST("/devices", makeCreateDeviceHandler(log, svc, v))
}
