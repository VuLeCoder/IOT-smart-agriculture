package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID         uuid.UUID `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	DeviceName string    `db:"device_name"`
	APIKey     string    `db:"api_key"`
	Location   string    `db:"location"`
	CreatedAt  time.Time `db:"created_at"`
}
