package middlewares

import (
	"IOT-Smart-Agriculture/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserAuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid authorization header",
			})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwtService.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user id in token",
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
