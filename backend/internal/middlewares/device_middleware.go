package middlewares

import (
	"IOT-Smart-Agriculture/internal/repositories"
	"IOT-Smart-Agriculture/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeviceAuthMiddleware(deviceRepo repositories.IDeviceRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		apiKey := c.GetHeader("IOT-API-Key")

		if apiKey == "" {
			response.Error(c, http.StatusUnauthorized, "Missing API Key", nil)
			c.Abort()
			return
		}

		deviceID, err := deviceRepo.VerifyDeviceByKey(c.Request.Context(), apiKey)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid API Key", err)
			c.Abort()
			return
		}

		c.Set("deviceID", deviceID)
		c.Next()
	}
}
