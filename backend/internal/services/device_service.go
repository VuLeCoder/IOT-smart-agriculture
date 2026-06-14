package services

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/models"
	"IOT-Smart-Agriculture/internal/repositories"
	"context"
	"time"

	"github.com/google/uuid"
)

type IDeviceService interface {
	CreateDevice(ctx context.Context, userID uuid.UUID, device dto.CreateDeviceRequest) (uuid.UUID, time.Time, error)
}

type deviceService struct {
	deviceRepo repositories.IDeviceRepository
}

func CreateNewDeviceService(deviceRepo repositories.IDeviceRepository) IDeviceService {
	return &deviceService{
		deviceRepo: deviceRepo,
	}
}

func (s *deviceService) CreateDevice(ctx context.Context, userID uuid.UUID, device dto.CreateDeviceRequest) (uuid.UUID, time.Time, error) {
	deviceModel := models.Device{
		ID:         uuid.New(),
		UserID:     userID,
		DeviceName: device.DeviceName,
		APIKey:     device.APIKey,
		Location:   device.Location,
		CreatedAt:  device.CreatedAt,
	}

	err := s.deviceRepo.CreateDevice(ctx, deviceModel)
	if err != nil {
		return uuid.Nil, time.Time{}, err
	}

	return deviceModel.ID, deviceModel.CreatedAt, nil
}
