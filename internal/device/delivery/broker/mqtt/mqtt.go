// Package mqtt provides impementation with MQTT.
package mqtt

import "fmt"

const (
	defaultQOSLevel = 1

	serverOrigin   = "api"
	firmwareOrigin = "firmware"
)

// HandlerFunc is a callback for event consuming.
//
// Parameters:
//   - []byte: Payload incoming from broker.
//
// Returns:
//   - error: if there is on error handling the message.
type HandlerFunc func(msg []byte) error

func buildTopic(deviceID, event, origin string) string {
	return fmt.Sprintf("devices/%s/event/%s/from/%s", deviceID, event, origin)
}
