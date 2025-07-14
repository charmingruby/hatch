package device

import (
	"github/charmingruby/pack/internal/device/delivery/rest"
	"github/charmingruby/pack/internal/device/repository/memory"
	"github/charmingruby/pack/internal/device/service"

	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
)

func New(log *logger.Logger, r *gin.Engine, val *validator.Validator) error {
	deviceRepo := memory.NewDeviceRepository()

	svc := service.New(&deviceRepo)

	rest.RegisterRoutes(log, r, svc, val)

	return nil
}
