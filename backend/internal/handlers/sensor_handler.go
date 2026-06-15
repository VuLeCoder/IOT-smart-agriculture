package handlers

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/services"
	"IOT-Smart-Agriculture/utils/response"
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
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	deviceID := c.MustGet("deviceID").(uuid.UUID)

	createAt, err := h.sensorService.SaveData(c.Request.Context(), deviceID, sensorDataReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to save sensor data", err)
		return
	}

	response.Success(c, http.StatusCreated, "Sensor data saved successfully", gin.H{
		"created_at": createAt,
	})
}

func (h *sensorHandler) GetData(c *gin.Context) {
	numStr := c.DefaultQuery("number", "1")

	number, err := strconv.Atoi(numStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Number must be an integer", err)
		return
	}

	id := c.Param("deviceID")
	deviceID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid device ID format", err)
		return
	}

	sensorData, err := h.sensorService.GetData(c.Request.Context(), deviceID, number)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to fetch sensor data", err)
		return
	}

	response.Success(c, http.StatusOK, "Sensor data retrieved successfully", gin.H{
		"data": sensorData,
	})
}
