package mqtt

import (
	"encoding/json"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DevicePublisher struct {
	cl  mqtt.Client
	log *logger.Logger
}

func NewDevicePublisher(log *logger.Logger, cl mqtt.Client) *DevicePublisher {
	return &DevicePublisher{
		cl:  cl,
		log: log,
	}
}

func (p *DevicePublisher) DispatchDeviceRegistered(msg broker.DeviceRegisteredMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	topic := buildTopic(msg.DeviceID, "registered", serverOrigin)

	token := p.cl.Publish(topic, defaultQOSLevel, true, payload)

	if err := token.Error(); err != nil {
		return err
	}

	p.log.Debug("message published", "topic", topic, "message", string(payload))

	return nil
}
