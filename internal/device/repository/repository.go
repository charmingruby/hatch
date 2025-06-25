package repository

import "github/charmingruby/habits/internal/device/model"

type DeviceRepository interface {
	Create(device model.Device) error
}
