package dto

import "time"

type CreateDeviceInput struct {
	TriggeredAt  time.Time `json:"triggered_at"  validate:"required,min=1"`
	HardwareID   string    `json:"hardware_id"   validate:"required,min=1"`
	HardwareType string    `json:"hardware_type" validate:"required,min=1"`
}
