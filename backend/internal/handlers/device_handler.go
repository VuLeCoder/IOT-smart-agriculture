package handlers

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IDeviceHandler interface {
	CreateDevice(c *gin.Context)
}

type deviceHandler struct {
	deviceService services.IDeviceService
}

func CreateNewDeviceHandler(deviceService services.IDeviceService) *deviceHandler {
	return &deviceHandler{
		deviceService: deviceService,
	}
}

func (h *deviceHandler) CreateDevice(c *gin.Context) {
	var deviceDataReq dto.CreateDeviceRequest

	if err := c.ShouldBindJSON(&deviceDataReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userID := c.MustGet("userID").(uuid.UUID)

	deviceID, createdAt, err := h.deviceService.CreateDevice(c.Request.Context(), userID, deviceDataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "create device successful",
		"device_id":  deviceID,
		"created_at": createdAt,
	})
}
