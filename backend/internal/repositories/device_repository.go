package repositories

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDeviceRepository interface {
	CreateDevice(ctx context.Context, device models.Device) error
	VerifyDeviceByKey(ctx context.Context, apiKey string) (uuid.UUID, error)
	GetDevicesByUser(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error)
}

type deviceRepo struct {
	db *pgxpool.Pool
}

func CreateNewDeviceRepo(db *pgxpool.Pool) *deviceRepo {
	return &deviceRepo{
		db: db,
	}
}

const (
	ADD_DEVICE_QUERY = `
		insert into devices (id, user_id, device_name, api_key, location, created_at)
		values ($1, $2, $3, $4, $5, $6)
	`

	VERIFY_DEVICE_QUERY = `
		select id
		from devices
		where api_key = $1
		limit 1
	`

	GET_DEVICE_BY_USER_QUERY = `
		select 
			id,
			user_id,
			api_key,
			device_name,
			location,
			created_at 
		from devices
		where user_id = $1
	`
)

func (r *deviceRepo) CreateDevice(ctx context.Context, device models.Device) error {
	_, err := r.db.Exec(
		ctx, ADD_DEVICE_QUERY,
		device.ID, device.UserID, device.DeviceName, device.APIKey, device.Location, device.CreatedAt,
	)
	return err
}

func (r *deviceRepo) VerifyDeviceByKey(ctx context.Context, apiKey string) (uuid.UUID, error) {
	var deviceID uuid.UUID

	err := r.db.QueryRow(ctx, VERIFY_DEVICE_QUERY, apiKey).Scan(&deviceID)
	if err != nil {
		return uuid.Nil, err
	}

	return deviceID, nil
}

func (r *deviceRepo) GetDevicesByUser(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error) {
	rows, err := r.db.Query(ctx, GET_DEVICE_BY_USER_QUERY, userID)
	if err != nil {
		return nil, err
	}

	var listDevices []dto.DeviceResponse
	for rows.Next() {
		var device dto.DeviceResponse

		err := rows.Scan(
			&device.ID, &device.UserID, &device.APIKey,
			&device.DeviceName, &device.Location, &device.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		listDevices = append(listDevices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listDevices, nil
}
