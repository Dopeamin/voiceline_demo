package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voiceline/backend/internal/interface/dto"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	version string
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		version: version,
	}
}

// HealthCheck returns the health status of the service
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dto.HealthResponseDTO{
		Status:  "healthy",
		Version: h.version,
	})
}

