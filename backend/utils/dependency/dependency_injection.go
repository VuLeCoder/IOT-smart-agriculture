package dependency

import (
	"IOT-Smart-Agriculture/internal/handlers"
	"IOT-Smart-Agriculture/internal/repositories"
	"IOT-Smart-Agriculture/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DI struct {
	DeviceRepo repositories.IDeviceRepository
	SensorRepo repositories.ISensorRepository

	DeviceService services.IDeviceService
	SensorService services.ISensorService

	DeviceHandler handlers.IDeviceHandler
	SensorHandler handlers.ISensorHandler
}

func CreateNewDI(db *pgxpool.Pool) *DI {

	deviceRepo := repositories.CreateNewDeviceRepo(db)
	sensorRepo := repositories.CreateNewSensorRepo(db)

	deviceService := services.CreateNewDeviceService(deviceRepo)
	sensorService := services.CreateNewSensorService(sensorRepo)

	deviceHandler := handlers.CreateNewDeviceHandler(deviceService)
	sensorHandler := handlers.CreateNewSensorHandler(sensorService)

	return &DI{
		DeviceRepo: deviceRepo,
		SensorRepo: sensorRepo,

		DeviceService: deviceService,
		SensorService: sensorService,

		DeviceHandler: deviceHandler,
		SensorHandler: sensorHandler,
	}
}
