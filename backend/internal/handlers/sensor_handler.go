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

// Save sensor data godoc
//
// @Summary Lưu dữ liệu cảm biến
// @Description Thiết bị gửi dữ liệu cảm biến (mưa, ánh sáng, độ ẩm, PH) về hệ thống
// @Tags Sensors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body dto.CreateSensorDataRequest true "Dữ liệu sensor"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/device/sensor-data [post]
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

// Get sensor data godoc
//
// @Summary Lấy lịch sử dữ liệu cảm biến
// @Description Người dùng lấy các bản ghi dữ liệu gần nhất của một thiết bị
// @Tags Sensors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param deviceID path string true "ID của thiết bị (UUID)"
// @Param number query int false "Số lượng bản ghi muốn lấy (mặc định 1)"
// @Success 200 {object} response.Response{data=[]dto.SensorDataResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/device/{deviceID}/sensor-data [get]
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
