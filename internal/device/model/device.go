package model

import (
	"github/charmingruby/gew/pkg/id"
	"time"
)

type Device struct {
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	ID           string    `json:"id"            db:"id"`
	HardwareID   string    `json:"hardware_id"   db:"hardware_id"`
	HardwareType string    `json:"hardware_type" db:"hardware_type"`
	IconURL      string    `json:"icon_url"      db:"icon_url"`
}

type DeviceInput struct {
	HardwareID   string
	HardwareType string
}

func NewDevice(in DeviceInput) Device {
	return Device{
		ID:           id.New(),
		HardwareID:   in.HardwareID,
		HardwareType: in.HardwareType,
		CreatedAt:    time.Now(),
	}
}
