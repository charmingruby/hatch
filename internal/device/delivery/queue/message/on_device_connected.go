package message

import "time"

type OnDeviceConnectedMessage struct {
	TriggeredAt  time.Time `json:"triggered_at"`
	HardwareID   string    `json:"hardware_id"`
	HardwareType string    `json:"hardware_type"`
}
