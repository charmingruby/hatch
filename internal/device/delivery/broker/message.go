package broker

import "time"

// DeviceBootedMessage is the payload to device booted event.
type DeviceBootedMessage struct {
	TriggeredAt  time.Time `json:"triggered_at"`
	HardwareID   string    `json:"hardware_id"`
	HardwareType string    `json:"hardware_type"`
}

// DeviceRegisteredMessage is the payload to device registered event.
type DeviceRegisteredMessage struct {
	DeviceID string `json:"device_id"`
}
