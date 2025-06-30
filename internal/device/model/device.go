package model

import (
	"github/charmingruby/pack/pkg/core"
	"time"
)

// Device represents the hardware data.
type Device struct {
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	ID           string    `json:"id"            db:"id"`
	HardwareID   string    `json:"hardware_id"   db:"hardware_id"`
	HardwareType string    `json:"hardware_type" db:"hardware_type"`
}

// DeviceInput is the params for Device constructor.
type DeviceInput struct {
	HardwareID   string
	HardwareType string
}

// NewDevice create an instance of Device
//
// Parameters:
//   - DeviceInput: the input data.
//
// Returns:
//   - Device: built Device structure.
func NewDevice(in DeviceInput) Device {
	return Device{
		ID:           core.NewID(),
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
		CreatedAt:    time.Now(),
	}
}
