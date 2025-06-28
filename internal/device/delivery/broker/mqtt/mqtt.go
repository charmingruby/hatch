package mqtt

import "fmt"

const (
	DEFAULT_QOS_LEVEL = 1

	SERVER_ORIGIN   = "api"
	FIRMWARE_ORIGIN = "firmware"
)

type HandlerFunc func(msg []byte) error

func buildTopic(deviceID, event, origin string) string {
	return fmt.Sprintf("devices/%s/event/%s/from/%s", deviceID, event, origin)
}
