package models

import (
	"time"

	"github.com/google/uuid"
)

type SensorData struct {
	ID           int64     `db:"id"`
	DeviceID     uuid.UUID `db:"device_id"`
	RainLevel    float64   `db:"rain_level"`
	Light        float64   `db:"light"`
	SoilMoisture float64   `db:"soil_moisture"`
	PH           float64   `db:"ph"`
	CreatedAt    time.Time `db:"created_at"`
}
