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

	deviceSub := mqtt.NewDeviceSubscriber(log, mqttCl)

	if err := deviceSub.SubscribeAll(); err != nil {
		return err
	}

	devicePub := mqtt.NewDevicePublisher(log, mqttCl)

	svc := service.New(deviceRepo, devicePub)

	rest.RegisterRoutes(log, r, svc, val)

	return nil
}
