package router

import (
	"github.com/gin-gonic/gin"

	"IOT-Smart-Agriculture/internal/middlewares"
	"IOT-Smart-Agriculture/utils/dependency"
)

func Setup(di *dependency.DI) *gin.Engine {
	r := gin.Default()
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

		webDevices.GET("/:deviceID/sensor-data", di.SensorHandler.GetData)
	}

	return r
}
