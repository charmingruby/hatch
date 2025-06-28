package mqtt

import "fmt"

const (
	SERVER_ORIGIN   = "api"
	FIRMWARE_ORIGIN = "firmware"
)

func buildTopic(deviceID, event, origin string) string {
	return fmt.Sprintf("devices/%s/event/%s/from/%s", deviceID, event, origin)
}
