package repository

import (
	"context"
	"github/charmingruby/pack/internal/device/model"
)

type DeviceRepository interface {
	Create(ctx context.Context, device model.Device) error
}
