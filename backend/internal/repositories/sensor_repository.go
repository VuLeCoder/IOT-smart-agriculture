package repositories

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ISensorRepository interface {
	SaveSensorData(ctx context.Context, sensorData models.SensorData) error
	GetLatestData(ctx context.Context, deviceID uuid.UUID) (dto.SensorDataResponse, error)
	GetHistoryData(ctx context.Context, deviceID uuid.UUID, number int) ([]dto.SensorDataResponse, error)
}

type sensorRepo struct {
	db *pgxpool.Pool
}

func CreateNewSensorRepo(db *pgxpool.Pool) ISensorRepository {
	return &sensorRepo{
		db: db,
	}
}

const (
	SAVE_SENSOR_DATA = `
		insert into sensor_data (id, device_id, rain_level, light, soil_moisture, ph, created_at)
		values (default, $1, $2, $3, $4, $5, $6)
	`

	GET_SENSOR_DATA = `
		select 
			rain_level, light, soil_moisture, ph, 
			created_at 
		from sensor_data
		where device_id = $1
		order by created_at desc
		limit 1
	`

	GET_HISTORY_SENSOR_DATA = `
		select 
			rain_level, light, soil_moisture, ph, 
			created_at 
		from sensor_data
		where device_id = $1
		order by created_at desc
		limit $2
	`
)

func (r *sensorRepo) SaveSensorData(ctx context.Context, sensorData models.SensorData) error {

	_, err := r.db.Exec(
		ctx, SAVE_SENSOR_DATA,
		sensorData.DeviceID, sensorData.RainLevel, sensorData.Light,
		sensorData.SoilMoisture, sensorData.PH, sensorData.CreatedAt,
	)
	return err
}

func (r *sensorRepo) GetLatestData(ctx context.Context, deviceID uuid.UUID) (dto.SensorDataResponse, error) {
	var data dto.SensorDataResponse
	err := r.db.QueryRow(ctx, GET_SENSOR_DATA, deviceID).Scan(
		&data.RainLevel, &data.Light, &data.SoilMoisture, &data.PH, &data.CreatedAt,
	)
	if err != nil {
		return dto.SensorDataResponse{}, err
	}
	return data, nil
}

func (r *sensorRepo) GetHistoryData(ctx context.Context, deviceID uuid.UUID, number int) ([]dto.SensorDataResponse, error) {
	rows, err := r.db.Query(ctx, GET_HISTORY_SENSOR_DATA, deviceID, number)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []dto.SensorDataResponse
	for rows.Next() {
		var data dto.SensorDataResponse

		err := rows.Scan(&data.RainLevel, &data.Light, &data.SoilMoisture, &data.PH, &data.CreatedAt)
		if err != nil {
			return nil, err
		}

		history = append(history, data)
	}

	return history, nil
}
