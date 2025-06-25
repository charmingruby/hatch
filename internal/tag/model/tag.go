package model

import (
	"time"
)

type Tag struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id"         db:"id"`
	SensorID  string    `json:"sensor_id"  db:"sensor_id"`
	IconURL   string    `json:"icon_url"   db:"icon_url"`
}
