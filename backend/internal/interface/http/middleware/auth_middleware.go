package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/interface/dto"
)

const (
	// UserIDKey is the key used to store user ID in the context
	UserIDKey = "userID"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
				Message: "Authorization header is required",
				Code:    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
				Message: "Invalid authorization header format",
				Code:    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		token := parts[1]

		// Validate token
		userID, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
				Message: "Invalid or expired token",
				Code:    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set(UserIDKey, userID)
		c.Next()
	}
}

// GetUserIDFromContext extracts user ID from the Gin context
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return uuid.Nil, false
	}

	id, ok := userID.(uuid.UUID)
	return id, ok
}

