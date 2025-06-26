package repository

import "github/charmingruby/pack/internal/device/model"

type DeviceRepository interface {
	Create(device model.Device) error
}
