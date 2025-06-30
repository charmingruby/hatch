package rest

import (
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register all the routes with handlers.
//
// Parameters:
//   - *logger.Logger: app logger.
//   - *gin.Engine: http handler.
//   - service.UseCase: application core logic.
//   - *validator.Validate: validator for payloads.
func RegisterRoutes(log *logger.Logger, r *gin.Engine, svc service.UseCase, v *validator.Validator) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	v1.POST("/devices", createDeviceHandler(log, svc, v))
}
