package postgres

import "github/charmingruby/habits/internal/device/model"

type DeviceRepo struct{}

func (r *DeviceRepo) Create(device model.Device) error {
	return nil
}
