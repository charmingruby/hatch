package memory

import (
	"context"
	"github/charmingruby/pack/internal/example/model"
)

type DeviceRepository struct {
	item []model.Device
}

func NewDeviceRepository() DeviceRepository {
	return DeviceRepository{
		item: make([]model.Device, 0),
	}
}

func (r *DeviceRepository) FindByHardwareIDAndType(_ context.Context, hwID, hwType string) (model.Device, error) {
	for _, i := range r.item {
		if i.HardwareID == hwID && i.HardwareType == hwType {
			return i, nil
		}
	}

	return model.Device{}, nil
}

func (r *DeviceRepository) Create(_ context.Context, device model.Device) error {
	r.item = append(r.item, device)
	return nil
}
