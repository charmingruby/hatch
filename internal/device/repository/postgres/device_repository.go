package postgres

import "github/charmingruby/pack/internal/device/model"

const ()

type DeviceRepo struct{}

func NewDeviceRepo() {}

func (r *DeviceRepo) Create(device model.Device) error {
	return nil
}

func (r *DeviceRepo) deviceQueries() {}
