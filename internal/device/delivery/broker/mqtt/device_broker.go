package mqtt

import (
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const QOS_LEVEL = 1

type HandlerFunc func(msg []byte) error

type DeviceBroker struct {
	cl       mqtt.Client
	log      *logger.Logger
	handlers map[string]HandlerFunc
}

func NewDeviceBroker(log *logger.Logger, cl mqtt.Client) *DeviceBroker {
	return &DeviceBroker{
		cl:       cl,
		log:      log,
		handlers: make(map[string]HandlerFunc),
	}
}

func (m *DeviceBroker) RegisterHandler(topic string, fn HandlerFunc) {
	m.handlers[topic] = fn
}

func (m *DeviceBroker) SubscribeAll() error {
	for topic := range m.handlers {
		token := m.cl.Subscribe(topic, QOS_LEVEL, m.dispatch)

		token.Wait()

		if err := token.Error(); err != nil {
			return err
		}
	}

	return nil
}

func (m *DeviceBroker) OnDeviceBooted(msg []byte) error {
	m.log.Debug("message received", "message", string(msg))
	return nil
}

func (c *DeviceBroker) dispatch(client mqtt.Client, msg mqtt.Message) {
	handler, ok := c.handlers[msg.Topic()]
	if !ok {
		return
	}

	if err := handler(msg.Payload()); err != nil {
		c.log.Warn("message processing error", "topic", msg.Topic(), "error", err)
		return
	}
}
