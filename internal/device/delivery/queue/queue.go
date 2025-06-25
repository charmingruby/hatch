package queue

import "github/charmingruby/habits/internal/device/delivery/queue/message"

type DeviceQueue interface {
	OnDeviceConnected(msg message.OnDeviceConnectedMessage) error
}
