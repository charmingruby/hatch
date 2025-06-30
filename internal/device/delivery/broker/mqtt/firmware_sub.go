package mqtt

import (
	"github/charmingruby/pack/pkg/broker/mqtt"
	"github/charmingruby/pack/pkg/telemetry/logger"

	mqttLib "github.com/eclipse/paho.mqtt.golang"
)

// FirmwareSubscriber holds the structure to listen to firmware events.
type FirmwareSubscriber struct {
	cl       mqttLib.Client
	log      *logger.Logger
	handlers map[string]HandlerFunc
}

// NewFirmwareSubscriber creates a FirmwareSubscriber instance.
//
// Parameters:
//   - *logger.Logger: logger instance
//   - mqttLib.Cllient: MQTT client connection
//
// Returns
//   - *FirmwareSubscriber: instance to start listening to events
func NewFirmwareSubscriber(log *logger.Logger, cl mqttLib.Client) *FirmwareSubscriber {
	return &FirmwareSubscriber{
		cl:       cl,
		log:      log,
		handlers: make(map[string]HandlerFunc),
	}
}

// SubscribeAll assign handlers and subscribe to the respective topics.
//
// Returns:
//   - error: if there is an error on subscribing to any topic of the assigned handlers
func (s *FirmwareSubscriber) SubscribeAll() error {
	s.assignHandlersToTopics()

	for topic := range s.handlers {
		token := s.cl.Subscribe(topic, defaultQOSLevel, s.handleMessage)

		token.Wait()

		if err := token.Error(); err != nil {
			return err
		}
	}

	return nil
}

// OnDeviceBooted listening to the firmware topic: "devices/+/event/+/from/firmware"
//
// Parameters:
//   - []byte: event payload
//
// Returns:
//   - error: if there is an error on logic
func (s *FirmwareSubscriber) OnDeviceBooted(msg []byte) error {
	s.log.Debug("message received", "message", string(msg))

	return nil
}

func (s *FirmwareSubscriber) handleMessage(_ mqttLib.Client, msg mqttLib.Message) {
	topic := msg.Topic()

	for pattern, handler := range s.handlers {
		matches := mqtt.TopicMatchesFilter(pattern, topic)

		if matches {
			if err := handler(msg.Payload()); err != nil {
				s.log.Warn("message processing error", "topic", topic, "error", err)
			}

			return
		}
	}
}

func (s *FirmwareSubscriber) assignHandlersToTopics() {
	s.handlers[buildTopic("+", "+", firmwareOrigin)] = s.OnDeviceBooted
}
