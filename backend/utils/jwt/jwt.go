package jwt

import (
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secretKey []byte
	expire    time.Duration
}

func CreateNewJWTService(secretKey string, expire time.Duration) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		expire:    expire,
	}
}

func (j *JWTService) GenerateJWT(userID uuid.UUID) (string, error) {
	claims := jwtv5.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(j.expire).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}
