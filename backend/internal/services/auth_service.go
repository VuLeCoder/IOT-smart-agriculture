package services

import (
	"IOT-Smart-Agriculture/internal/dto"
	"IOT-Smart-Agriculture/internal/models"
	"IOT-Smart-Agriculture/internal/repositories"
	"IOT-Smart-Agriculture/utils/crypto"
	"IOT-Smart-Agriculture/utils/jwt"
	"context"
	"time"

	"github.com/google/uuid"
)

type IAuthService interface {
	Register(ctx context.Context, user dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(ctx context.Context, userLogin dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	userRepo   repositories.IUserRepository
	jwtService *jwt.JWTService
}

func CreateNewAuthService(userRepo repositories.IUserRepository, jwtService *jwt.JWTService) IAuthService {
	return &authService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *authService) Register(ctx context.Context, user dto.RegisterRequest) (dto.RegisterResponse, error) {
	passwordHash, err := crypto.HashPassword(user.Password)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	var userModel = models.User{
		ID:           uuid.New(),
		Email:        user.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	err = s.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	userResponse := dto.RegisterResponse{
		ID:        userModel.ID,
		CreatedAt: userModel.CreatedAt,
	}

	return userResponse, nil
}

func (s *authService) Login(ctx context.Context, userLogin dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, userLogin.Email)
	if err != nil {
		return dto.LoginResponse{}, ErrInvalidCredentials
	}

	if !crypto.CheckPassword(userLogin.Password, user.PasswordHash) {
		return dto.LoginResponse{}, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateJWT(user.ID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	loginResponse := dto.LoginResponse{
		Token: token,
		ID:    user.ID,
		Email: user.Email,
	}
	return loginResponse, nil
}
