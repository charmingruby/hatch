package example

import (
	"github/charmingruby/pack/internal/example/delivery/http/rest"
	"github/charmingruby/pack/internal/example/repository/memory"
	"github/charmingruby/pack/internal/example/service"

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
