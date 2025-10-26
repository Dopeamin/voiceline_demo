package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/interface/dto"
	"github.com/voiceline/backend/internal/interface/mappers"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *services.AuthService
	userMapper  *mappers.UserMapper
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *services.AuthService, userMapper *mappers.UserMapper) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userMapper:  userMapper,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Message: err.Error(),
			Code:    "INVALID_REQUEST",
		})
		return
	}

	output, err := h.authService.Register(c.Request.Context(), services.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})

	if err != nil {
		statusCode := http.StatusInternalServerError
		code := "INTERNAL_ERROR"

		if err == services.ErrUserAlreadyExists {
			statusCode = http.StatusConflict
			code = "USER_EXISTS"
		}

		c.JSON(statusCode, dto.ErrorDTO{
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	response := h.userMapper.ToAuthResponseDTO(output.Token, output.User)
	c.JSON(http.StatusCreated, response)
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Message: err.Error(),
			Code:    "INVALID_REQUEST",
		})
		return
	}

	output, err := h.authService.Login(c.Request.Context(), services.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		statusCode := http.StatusInternalServerError
		code := "INTERNAL_ERROR"

		if err == services.ErrInvalidCredentials {
			statusCode = http.StatusUnauthorized
			code = "INVALID_CREDENTIALS"
		}

		c.JSON(statusCode, dto.ErrorDTO{
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	response := h.userMapper.ToAuthResponseDTO(output.Token, output.User)
	c.JSON(http.StatusOK, response)
}
