package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/interface/dto"
	"github.com/voiceline/backend/internal/interface/http/middleware"
	"github.com/voiceline/backend/internal/interface/mappers"
)

// TranscriptionHandler handles transcription requests
type TranscriptionHandler struct {
	transcriptionService *services.TranscriptionService
	transcriptionMapper  *mappers.TranscriptionMapper
}

// NewTranscriptionHandler creates a new TranscriptionHandler
func NewTranscriptionHandler(
	transcriptionService *services.TranscriptionService,
	transcriptionMapper *mappers.TranscriptionMapper,
) *TranscriptionHandler {
	return &TranscriptionHandler{
		transcriptionService: transcriptionService,
		transcriptionMapper:  transcriptionMapper,
	}
}

// TranscribeAudio handles audio transcription
func (h *TranscriptionHandler) TranscribeAudio(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
			Message: "Unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	file, err := c.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Message: "Audio file is required",
			Code:    "INVALID_REQUEST",
		})
		return
	}

	// Open the uploaded file
	audioFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Message: "Failed to read audio file",
			Code:    "INTERNAL_ERROR",
		})
		return
	}
	defer audioFile.Close()

	// Transcribe audio
	transcription, err := h.transcriptionService.Transcribe(c.Request.Context(), services.TranscribeAudioInput{
		UserID: userID,
		Audio:  audioFile,
	})

	if err != nil {
		statusCode := http.StatusInternalServerError
		code := "TRANSCRIPTION_FAILED"

		// Check if it's a configuration error
		if err.Error() == "OpenAI service is not configured. Please set OPENAI_API_KEY environment variable" {
			statusCode = http.StatusServiceUnavailable
			code = "SERVICE_NOT_CONFIGURED"
		}

		c.JSON(statusCode, dto.ErrorDTO{
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	response := h.transcriptionMapper.ToDTO(transcription)
	c.JSON(http.StatusOK, response)
}

// GetTranscriptions gets all transcriptions for the authenticated user
func (h *TranscriptionHandler) GetTranscriptions(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
			Message: "Unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	transcriptions, err := h.transcriptionService.GetUserTranscriptions(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Message: err.Error(),
			Code:    "INTERNAL_ERROR",
		})
		return
	}

	response := h.transcriptionMapper.ToDTOs(transcriptions)
	c.JSON(http.StatusOK, response)
}

// GetTranscription gets a specific transcription by ID
func (h *TranscriptionHandler) GetTranscription(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.ErrorDTO{
			Message: "Unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Message: "Invalid transcription ID",
			Code:    "INVALID_REQUEST",
		})
		return
	}

	transcription, err := h.transcriptionService.GetTranscription(c.Request.Context(), id, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		code := "INTERNAL_ERROR"

		if err == services.ErrTranscriptionNotFound {
			statusCode = http.StatusNotFound
			code = "NOT_FOUND"
		} else if err == services.ErrUnauthorizedAccess {
			statusCode = http.StatusForbidden
			code = "FORBIDDEN"
		}

		c.JSON(statusCode, dto.ErrorDTO{
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	response := h.transcriptionMapper.ToDTO(transcription)
	c.JSON(http.StatusOK, response)
}
