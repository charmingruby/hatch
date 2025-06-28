package broker

type DeviceSubscriber interface {
	OnDeviceBooted(msg []byte) error
}

type DevicePublisher interface {
	DispatchDeviceRegistered(msg DeviceRegisteredMessage) error
}
