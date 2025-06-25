package message

import "time"

type DispatchDeviceCreatedMessage struct {
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	ID           string    `json:"id"            db:"id"`
	HardwareID   string    `json:"hardware_id"   db:"hardware_id"`
	HardwareType string    `json:"hardware_type" db:"hardware_type"`
}
