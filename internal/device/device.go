package device

import (
	"github/charmingruby/pack/internal/device/delivery/rest"
	"github/charmingruby/pack/internal/device/repository/postgres"
	"github/charmingruby/pack/internal/device/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(r *gin.Engine, db *sqlx.DB) error {
	deviceRepo, err := postgres.NewDeviceRepo(db)
	if err != nil {
		return err
	}

	svc := service.New(deviceRepo)

	rest.RegisterRoutes(r, svc)

	return nil
}
