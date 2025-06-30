// Package repository provides the repositories contracts.
package repository

import (
	"context"
	"github/charmingruby/pack/internal/device/model"
)

// DeviceRepository is the contract to handle persistece of model.Device.
type DeviceRepository interface {
	FindByHardwareIDAndType(ctx context.Context, hwID, hwType string) (model.Device, error)
	Create(ctx context.Context, device model.Device) error
}
