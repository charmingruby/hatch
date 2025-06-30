// Package broker provides contracts for messaging.
package broker

// FirmwareSubscriber is the contract for consuming menssages from firmware.
type FirmwareSubscriber interface {
	OnDeviceBooted(msg []byte) error
}

// FirmwarePublisher is the contract for publishing menssages from firmware.
type FirmwarePublisher interface {
	DispatchDeviceRegistered(msg DeviceRegisteredMessage) error
}
