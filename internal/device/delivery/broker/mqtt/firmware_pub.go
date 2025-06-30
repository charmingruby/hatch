package mqtt

import (
	"encoding/json"
	"github/charmingruby/pack/internal/device/delivery/broker"
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// FirmwarePublisher holds the structure to publish to firmware events.
type FirmwarePublisher struct {
	cl  mqtt.Client
	log *logger.Logger
}

// NewFirmwarePublisher creates a FirmwarePublisher instance.
//
// Parameters:
//   - *logger.Logger: logger instance.
//   - mqttLib.Cllient: MQTT client connection.
//
// Returns
//   - *FirmwarePublisher: instance to start publishing to events.
func NewFirmwarePublisher(log *logger.Logger, cl mqtt.Client) *FirmwarePublisher {
	return &FirmwarePublisher{
		cl:  cl,
		log: log,
	}
}

// DispatchDeviceRegistered publishes a message to the device topic: "devices/{device_id}/event/registered/from/api".
//
// Parameters:
//   - broker.DeviceRegisteredMessage: message for the respective event.
//
// Returns
//   - error: if there is an error on serializing or on publishing.
func (p *FirmwarePublisher) DispatchDeviceRegistered(msg broker.DeviceRegisteredMessage) error {
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
