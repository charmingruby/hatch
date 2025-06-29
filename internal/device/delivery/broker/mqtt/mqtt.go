package mqtt

import "fmt"

const (
	defaultQOSLevel = 1

	serverOrigin   = "api"
	firmwareOrigin = "firmware"
)

type HandlerFunc func(msg []byte) error

func buildTopic(deviceID, event, origin string) string {
	return fmt.Sprintf("devices/%s/event/%s/from/%s", deviceID, event, origin)
}
