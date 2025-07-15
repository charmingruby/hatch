package repository

import (
	"context"
	"github/charmingruby/pack/internal/example/model"
)

type DeviceRepository interface {
	FindByHardwareIDAndType(ctx context.Context, hwID, hwType string) (model.Device, error)
	Create(ctx context.Context, device model.Device) error
}
