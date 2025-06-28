package mqtt

import (
	"github/charmingruby/pack/pkg/broker/mqtt"
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqttLib "github.com/eclipse/paho.mqtt.golang"
)

var (
	onDeviceBootedTopic = buildTopic("+", "+", SERVER_ORIGIN)
)

type DeviceSubscriber struct {
	cl       mqttLib.Client
	log      *logger.Logger
	handlers map[string]HandlerFunc
}

func NewDeviceSubscriber(log *logger.Logger, cl mqttLib.Client) *DeviceSubscriber {
	return &DeviceSubscriber{
		cl:       cl,
		log:      log,
		handlers: make(map[string]HandlerFunc),
	}
}

func (m *DeviceSubscriber) SubscribeAll() error {
	m.assignHandlersToTopics()

	for topic := range m.handlers {
		token := m.cl.Subscribe(topic, DEFAULT_QOS_LEVEL, m.handleMessage)

		token.Wait()

		if err := token.Error(); err != nil {
			return err
		}
	}

	return nil
}

func (m *DeviceSubscriber) OnDeviceBooted(msg []byte) error {
	m.log.Debug("message received", "topic", onDeviceBootedTopic, "message", string(msg))

	return nil
}

func (c *DeviceSubscriber) handleMessage(client mqttLib.Client, msg mqttLib.Message) {
	topic := msg.Topic()

	for pattern, handler := range c.handlers {
		matches := mqtt.TopicMatchesFilter(pattern, topic)

		if matches {
			if err := handler(msg.Payload()); err != nil {
				c.log.Warn("message processing error", "topic", topic, "error", err)
			}

			return
		}
	}
}

func (m *DeviceSubscriber) assignHandlersToTopics() {
	m.handlers[onDeviceBootedTopic] = m.OnDeviceBooted
}
