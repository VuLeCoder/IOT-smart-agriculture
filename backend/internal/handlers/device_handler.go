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
	GetDevices(c *gin.Context)
	GetDeviceByID(c *gin.Context)
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
		return
	}

	userID := c.MustGet("userID").(uuid.UUID)

	deviceID, createdAt, err := h.deviceService.CreateDevice(c.Request.Context(), userID, deviceDataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "create device successful",
		"device_id":  deviceID,
		"created_at": createdAt,
	})
}

func (h *deviceHandler) GetDevices(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	listDevices, err := h.deviceService.GetDevices(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get devices successful",
		"devices": listDevices,
	})
}

func (h *deviceHandler) GetDeviceByID(c *gin.Context) {

	userID := c.MustGet("userID").(uuid.UUID)

	id := c.Param("deviceID")
	deviceID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	device, err := h.deviceService.GetDeviceByID(c.Request.Context(), userID, deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get device successful",
		"device":  device,
	})
}
