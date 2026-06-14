package dependency

import (
	"IOT-Smart-Agriculture/internal/config"
	"IOT-Smart-Agriculture/internal/handlers"
	"IOT-Smart-Agriculture/internal/repositories"
	"IOT-Smart-Agriculture/internal/services"
	"IOT-Smart-Agriculture/utils/jwt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DI struct {
	DeviceRepo repositories.IDeviceRepository
	SensorRepo repositories.ISensorRepository
	UserRepo   repositories.IUserRepository

	AuthService   services.IAuthService
	DeviceService services.IDeviceService
	SensorService services.ISensorService

	AuthHandler   handlers.IAuthHandler
	DeviceHandler handlers.IDeviceHandler
	SensorHandler handlers.ISensorHandler
}

func CreateNewDI(db *pgxpool.Pool, cfg config.Config) *DI {

	jwtService := jwt.CreateNewJWTService(cfg.JWTSecretKey, time.Duration(cfg.JWTExpireHours))

	deviceRepo := repositories.CreateNewDeviceRepo(db)
	sensorRepo := repositories.CreateNewSensorRepo(db)
	userRepo := repositories.CreateNewUserRepo(db)

	authService := services.CreateNewAuthService(userRepo, jwtService)
	deviceService := services.CreateNewDeviceService(deviceRepo)
	sensorService := services.CreateNewSensorService(sensorRepo)

	authHandler := handlers.CreateNewAuthHandler(authService)
	deviceHandler := handlers.CreateNewDeviceHandler(deviceService)
	sensorHandler := handlers.CreateNewSensorHandler(sensorService)

	return &DI{
		DeviceRepo: deviceRepo,
		SensorRepo: sensorRepo,
		UserRepo:   userRepo,

		AuthService:   authService,
		DeviceService: deviceService,
		SensorService: sensorService,

		AuthHandler:   authHandler,
		DeviceHandler: deviceHandler,
		SensorHandler: sensorHandler,
	}
}
