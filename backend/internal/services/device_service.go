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
	GetDevices(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error)
	GetDeviceByID(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID) (dto.DeviceResponse, error)
	UpdateDevice(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID, req dto.UpdateDeviceRequest) error
	DeleteDevice(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID) error
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
		CreatedAt:  time.Now(),
	}

	err := s.deviceRepo.CreateDevice(ctx, deviceModel)
	if err != nil {
		return uuid.Nil, time.Time{}, err
	}

	return deviceModel.ID, deviceModel.CreatedAt, nil
}

func (s *deviceService) GetDevices(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error) {
	listDevices, err := s.deviceRepo.GetDevicesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return listDevices, nil
}

func (s *deviceService) GetDeviceByID(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID) (dto.DeviceResponse, error) {
	device, err := s.deviceRepo.GetDeviceByID(ctx, userID, deviceID)
	if err != nil {
		return dto.DeviceResponse{}, err
	}

	return device, nil
}

func (s *deviceService) UpdateDevice(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID, req dto.UpdateDeviceRequest) error {
	return s.deviceRepo.UpdateDevice(ctx, userID, deviceID, req)
}

func (s *deviceService) DeleteDevice(ctx context.Context, userID uuid.UUID, deviceID uuid.UUID) error {
	return s.deviceRepo.DeleteDevice(ctx, userID, deviceID)
}
