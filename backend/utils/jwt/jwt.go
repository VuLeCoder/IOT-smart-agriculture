package jwt

import (
	"errors"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secretKey []byte
	expire    time.Duration
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwtv5.RegisteredClaims
}

func CreateNewJWTService(secretKey string, expire time.Duration) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		expire:    expire,
	}
}

func (j *JWTService) GenerateJWT(userID uuid.UUID) (string, error) {
	claims := UserClaims{
		UserID: userID.String(),
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(j.expire)),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
		},
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *JWTService) ValidateJWT(tokenStr string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwtv5.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwtv5.Token) (interface{}, error) {
			return j.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
