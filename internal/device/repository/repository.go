package repository

import "github/charmingruby/gew/internal/device/model"

type DeviceRepository interface {
	Create(device model.Device) error
}
