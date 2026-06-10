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

	SensorService services.ISensorService

	SensorHandler handlers.ISensorHandler
}

func CreateNewDI(db *pgxpool.Pool) *DI {

	deviceRepo := repositories.CreateNewDeviceRepo(db)
	sensorRepo := repositories.CreateNewSensorRepo(db)

	sensorService := services.CreateNewSensorService(sensorRepo)

	sensorHandler := handlers.CreateNewSensorHandler(sensorService)

	return &DI{
		DeviceRepo: deviceRepo,
		SensorRepo: sensorRepo,

		SensorService: sensorService,

		SensorHandler: sensorHandler,
	}
}
