package handlers

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/services"
	"IOT-Smart-Agriculture/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	authService services.IAuthService
}

func CreateNewAuthHandler(authService services.IAuthService) IAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

// Register godoc
//
// @Summary Đăng ký tài khoản
// @Description Người dùng đăng ký tài khoản mới bằng email và mật khẩu
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Thông tin đăng ký"
// @Success 201 {object} response.Response{data=dto.RegisterResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var registerReq dto.RegisterRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	registerResponse, err := h.authService.Register(c.Request.Context(), registerReq)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Registration failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Registration successful", registerResponse)
}

// Login godoc
//
// @Summary Đăng nhập
// @Description Người dùng đăng nhập để lấy JWT Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Thông tin đăng nhập"
// @Success 200 {object} response.Response{data=dto.LoginResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /IOT-api/web/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var loginReq dto.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	loginResponse, err := h.authService.Login(c.Request.Context(), loginReq)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Login failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Login successful", loginResponse)
}
