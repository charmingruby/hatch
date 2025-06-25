package postgres

import "github/charmingruby/gew/internal/device/model"

type DeviceRepo struct{}

func (r *DeviceRepo) Create(device model.Device) error {
	return nil
}
