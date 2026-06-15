package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"IOT-Smart-Agriculture/internal/middlewares"
	"IOT-Smart-Agriculture/utils/dependency"
)

func Setup(di *dependency.DI) *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/IOT-api")

	device := api.Group("/device")
	device.Use(middlewares.DeviceAuthMiddleware(di.DeviceRepo))
	{
		device.POST("/sensor-data", di.SensorHandler.SaveData)
	}

	web := api.Group("/web")
	web.POST("/register", di.AuthHandler.Register)
	web.POST("/login", di.AuthHandler.Login)

	web.Use(middlewares.UserAuthMiddleware(&di.JWTService))
	webDevices := web.Group("/devices")
	{
		webDevices.POST("", di.DeviceHandler.CreateDevice)
		webDevices.GET("", di.DeviceHandler.GetDevices)

		webDevices.GET("/:deviceID", di.DeviceHandler.GetDeviceByID)
		webDevices.PATCH("/:deviceID", di.DeviceHandler.UpdateDeviceByID)
		webDevices.DELETE("/:deviceID", di.DeviceHandler.DeleteDeviceByID)

		webDevices.GET("/:deviceID/sensor-data", di.SensorHandler.GetData)

	}

	return r
}
