package handlers

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ISensorHandler interface {
	SaveData(c *gin.Context)
	GetData(c *gin.Context)
}

type sensorHandler struct {
	sensorService services.ISensorService
}

func CreateNewSensorHandler(sensorService services.ISensorService) *sensorHandler {
	return &sensorHandler{
		sensorService: sensorService,
	}
}

func (h *sensorHandler) SaveData(c *gin.Context) {
	var sensorDataReq dto.CreateSensorDataRequest

	if err := c.ShouldBindJSON(&sensorDataReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	deviceID := c.MustGet("deviceID").(uuid.UUID)

	createAt, err := h.sensorService.SaveData(c.Request.Context(), deviceID, sensorDataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "sensor data saved",
		"created_at": createAt,
	})
}

func (h *sensorHandler) GetData(c *gin.Context) {
	numStr := c.DefaultQuery("number", "1")

	number, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "number must be integer",
		})
		return
	}

	deviceID := c.MustGet("deviceID").(uuid.UUID)

	sensorData, err := h.sensorService.GetData(c.Request.Context(), deviceID, number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get data successful",
		"data":    sensorData,
	})
}
