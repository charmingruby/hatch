package mqtt

import (
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DeviceSubscriber struct {
	cl       mqtt.Client
	log      *logger.Logger
	handlers map[string]HandlerFunc
}

func NewDeviceSubscriber(log *logger.Logger, cl mqtt.Client) *DeviceSubscriber {
	return &DeviceSubscriber{
		cl:       cl,
		log:      log,
		handlers: make(map[string]HandlerFunc),
	}
}

func (m *DeviceSubscriber) RegisterHandler(topic string, fn HandlerFunc) {
	m.handlers[topic] = fn
}

func (m *DeviceSubscriber) SubscribeAll() error {
	for topic := range m.handlers {
		token := m.cl.Subscribe(topic, QOS_LEVEL, m.handleMessage)

		token.Wait()

		if err := token.Error(); err != nil {
			return err
		}
	}

	return nil
}

func (m *DeviceSubscriber) OnDeviceBooted(msg []byte) error {
	m.log.Debug("message received", "message", string(msg))

	return nil
}

func (c *DeviceSubscriber) handleMessage(client mqtt.Client, msg mqtt.Message) {
	handler, ok := c.handlers[msg.Topic()]
	if !ok {
		return
	}

	if err := handler(msg.Payload()); err != nil {
		c.log.Warn("message processing error", "topic", msg.Topic(), "error", err)
		return
	}
}
