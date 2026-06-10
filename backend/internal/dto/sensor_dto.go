package dto

import "time"

type CreateSensorDataRequest struct {
	RainLevel    float64 `json:"rain_level"`
	Light        float64 `json:"light"`
	SoilMoisture float64 `json:"soil_moisture"`
	PH           float64 `json:"ph"`
}

type SensorDataResponse struct {
	RainLevel    float64   `json:"rain_level"`
	Light        float64   `json:"light"`
	SoilMoisture float64   `json:"soil_moisture"`
	PH           float64   `json:"ph"`
	CreatedAt    time.Time `json:"created_at"`
}
