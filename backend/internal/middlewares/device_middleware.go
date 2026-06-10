package middlewares

import (
	"IOT-Smart-Agriculture/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeviceAuthMiddleware(deviceRepo repositories.IDeviceRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		apiKey := c.GetHeader("IOT-API-Key")

		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing api key",
			})
			return
		}

		deviceID, err := deviceRepo.VerifyDeviceByKey(c.Request.Context(), apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api key",
			})
			return
		}

		c.Set("deviceID", deviceID)
		c.Next()
	}
}
