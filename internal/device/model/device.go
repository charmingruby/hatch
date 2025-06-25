package model

import (
	"time"
)

type Device struct {
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	ID           string    `json:"id"            db:"id"`
	HardwareID   string    `json:"hardware_id"   db:"hardware_id"`
	HardwareType string    `json:"hardware_type" db:"hardware_type"`
	IconURL      string    `json:"icon_url"      db:"icon_url"`
}
