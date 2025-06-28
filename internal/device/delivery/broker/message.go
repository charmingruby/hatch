package broker

import "time"

type DeviceBootedMessage struct {
	TriggeredAt  time.Time `json:"triggered_at"`
	HardwareID   string    `json:"hardware_id"`
	HardwareType string    `json:"hardware_type"`
}

type DeviceRegisteredMessage struct {
	DeviceID string `json:"device_id"`
}
