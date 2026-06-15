package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	APIKey     string    `json:"api_key" binding:"required"`
	DeviceName string    `json:"device_name" binding:"required"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
}

type DeviceResponse struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	APIKey     string    `json:"api_key"`
	DeviceName string    `json:"device_name"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
}
