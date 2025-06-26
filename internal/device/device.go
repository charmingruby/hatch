package device

import (
	"github/charmingruby/pack/internal/device/delivery/rest"
	"github/charmingruby/pack/internal/device/repository/postgres"
	"github/charmingruby/pack/internal/device/service"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(r *gin.Engine, db *sqlx.DB, val *validator.Validator) error {
	deviceRepo, err := postgres.NewDeviceRepo(db)
	if err != nil {
		return err
	}

	svc := service.New(deviceRepo)

	rest.RegisterRoutes(r, svc, val)

	return nil
}
