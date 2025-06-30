package device

import (
	"github/charmingruby/pack/internal/device/delivery/broker/mqtt"
	"github/charmingruby/pack/internal/device/delivery/rest"
	"github/charmingruby/pack/internal/device/repository/postgres"
	"github/charmingruby/pack/internal/device/service"

	mqttLib "github.com/eclipse/paho.mqtt.golang"

	"github/charmingruby/pack/pkg/telemetry/logger"
	"github/charmingruby/pack/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(log *logger.Logger, mqttCl mqttLib.Client, r *gin.Engine, db *sqlx.DB, val *validator.Validator) error {
	deviceRepo, err := postgres.NewDeviceRepo(db)
	if err != nil {
		return err
	}

	firmwareSub := mqtt.NewFirmwareSubscriber(log, mqttCl)

	if err := firmwareSub.SubscribeAll(); err != nil {
		return err
	}

	firmwarePub := mqtt.NewFirmwarePublisher(log, mqttCl)

	svc := service.New(deviceRepo, firmwarePub)

	rest.RegisterRoutes(log, r, svc, val)

	return nil
}
