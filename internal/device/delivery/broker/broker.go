package broker

type DeviceBroker interface {
	OnDeviceBooted(msg []byte) error
}
