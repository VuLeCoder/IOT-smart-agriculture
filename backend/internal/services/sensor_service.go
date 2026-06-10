package services

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/models"
	"IOT-Smart-Agriculture/internal/repositories"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidNumber = errors.New("number must be between 1 and 100")
)

type ISensorService interface {
	SaveData(ctx context.Context, deviceID uuid.UUID, sensorData dto.SensorDataRequest) (time.Time, error)
	GetData(ctx context.Context, deviceID uuid.UUID, number int) ([]dto.SensorDataResponse, error)
}

type sensorService struct {
	sensorRepo repositories.ISensorRepository
}

func CreateNewSensorService(sensorRepo repositories.ISensorRepository) *sensorService {
	return &sensorService{
		sensorRepo: sensorRepo,
	}
}

func (s *sensorService) SaveData(ctx context.Context, deviceID uuid.UUID, sensorData dto.SensorDataRequest) (time.Time, error) {
	sensorDataModel := models.SensorData{
		DeviceID:     deviceID,
		RainLevel:    sensorData.RainLevel,
		Light:        sensorData.Light,
		SoilMoisture: sensorData.SoilMoisture,
		PH:           sensorData.PH,
		CreatedAt:    time.Now(),
	}

	err := s.sensorRepo.SaveSensorData(ctx, sensorDataModel)
	if err != nil {
		return time.Time{}, err
	}

	return sensorDataModel.CreatedAt, nil
}

func (s *sensorService) GetData(ctx context.Context, deviceID uuid.UUID, number int) ([]dto.SensorDataResponse, error) {
	if number < 1 || number > 100 {
		return nil, ErrInvalidNumber
	}

	return s.sensorRepo.GetLatestData(ctx, deviceID, number)
}
