package middlewares

import (
	"IOT-Smart-Agriculture/utils/jwt"
	"IOT-Smart-Agriculture/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserAuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "Missing or invalid authorization header", nil)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwtService.ValidateJWT(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token", err)
			c.Abort()
			return
		}

		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid user ID in token", err)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
