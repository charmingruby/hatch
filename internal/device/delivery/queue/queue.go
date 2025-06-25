package queue

type DeviceQueue interface {
	OnDeviceConnected(msg OnDeviceConnectedMessage) error
}
