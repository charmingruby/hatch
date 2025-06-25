package queue

type DeviceQueue interface {
	OnDeviceBooted(msg OnDeviceBootedMessage) error
}
