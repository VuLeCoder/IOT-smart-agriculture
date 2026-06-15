package handlers

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/services"
	"IOT-Smart-Agriculture/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IDeviceHandler interface {
	CreateDevice(c *gin.Context)
	GetDevices(c *gin.Context)
	GetDeviceByID(c *gin.Context)
	UpdateDeviceByID(c *gin.Context)
	DeleteDeviceByID(c *gin.Context)
}

type deviceHandler struct {
	deviceService services.IDeviceService
}

func CreateNewDeviceHandler(deviceService services.IDeviceService) *deviceHandler {
	return &deviceHandler{
		deviceService: deviceService,
	}
}

// Create device godoc
//
// @Summary Thêm thiết bị mới
// @Description Người dùng thêm thiết bị (cảm biến) mới vào hệ thống
// @Tags Devices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateDeviceRequest true "Thông tin thiết bị mới"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/devices [post]
func (h *deviceHandler) CreateDevice(c *gin.Context) {
	var deviceDataReq dto.CreateDeviceRequest

	if err := c.ShouldBindJSON(&deviceDataReq); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	userID := c.MustGet("userID").(uuid.UUID)

	deviceID, createdAt, err := h.deviceService.CreateDevice(c.Request.Context(), userID, deviceDataReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create device", err)
		return
	}

	response.Success(c, http.StatusCreated, "Device created successfully", gin.H{
		"device_id":  deviceID,
		"created_at": createdAt,
	})
}

// Get devices godoc
//
// @Summary Danh sách thiết bị
// @Description Lấy danh sách tất cả thiết bị của người dùng hiện tại
// @Tags Devices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]dto.DeviceResponse}
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/devices [get]
func (h *deviceHandler) GetDevices(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	listDevices, err := h.deviceService.GetDevices(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch devices", err)
		return
	}

	response.Success(c, http.StatusOK, "Devices retrieved successfully", gin.H{
		"devices": listDevices,
	})
}

// Get device by ID godoc
//
// @Summary Chi tiết thiết bị
// @Description Lấy thông tin chi tiết của một thiết bị theo ID
// @Tags Devices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param deviceID path string true "ID của thiết bị (UUID)"
// @Success 200 {object} response.Response{data=dto.DeviceResponse}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/devices/{deviceID} [get]
func (h *deviceHandler) GetDeviceByID(c *gin.Context) {

	userID := c.MustGet("userID").(uuid.UUID)

	id := c.Param("deviceID")
	deviceID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid device ID format", err)
		return
	}

	device, err := h.deviceService.GetDeviceByID(c.Request.Context(), userID, deviceID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch device details", err)
		return
	}

	response.Success(c, http.StatusOK, "Device details retrieved successfully", gin.H{
		"device": device,
	})
}

// Update device godoc
//
// @Summary Cập nhật thiết bị
// @Description Cập nhật thông tin của một thiết bị (chỉ cập nhật các trường được gửi)
// @Tags Devices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param deviceID path string true "ID của thiết bị (UUID)"
// @Param request body dto.UpdateDeviceRequest true "Thông tin cập nhật"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/devices/{deviceID} [patch]
func (h *deviceHandler) UpdateDeviceByID(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	id := c.Param("deviceID")
	deviceID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid device ID format", err)
		return
	}

	var updateReq dto.UpdateDeviceRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	err = h.deviceService.UpdateDevice(c.Request.Context(), userID, deviceID, updateReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update device", err)
		return
	}

	response.Success(c, http.StatusOK, "Device updated successfully", nil)
}

// Delete device godoc
//
// @Summary Xóa thiết bị
// @Description Xóa một thiết bị khỏi hệ thống
// @Tags Devices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param deviceID path string true "ID của thiết bị (UUID)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/devices/{deviceID} [delete]
func (h *deviceHandler) DeleteDeviceByID(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	id := c.Param("deviceID")
	deviceID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid device ID format", err)
		return
	}

	err = h.deviceService.DeleteDevice(c.Request.Context(), userID, deviceID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete device", err)
		return
	}

	response.Success(c, http.StatusOK, "Device deleted successfully", nil)
}
