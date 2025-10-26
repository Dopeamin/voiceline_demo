package mappers

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/voiceline/backend/internal/domain/entities"
	"github.com/voiceline/backend/internal/interface/mappers"
)

var NewTranscriptionMapper = mappers.NewTranscriptionMapper

func TestTranscriptionMapper_ToDTO(t *testing.T) {
	mapper := NewTranscriptionMapper()

	t.Run("Convert valid transcription", func(t *testing.T) {
		transcription := &entities.Transcription{
			ID:        uuid.New(),
			UserID:    uuid.New(),
			Text:      "Test transcription",
			Status:    entities.StatusCompleted,
			Duration:  10.5,
			CreatedAt: time.Now(),
		}

		dto := mapper.ToDTO(transcription)

		assert.NotNil(t, dto)
		assert.Equal(t, transcription.ID.String(), dto.ID)
		assert.Equal(t, transcription.Text, dto.Text)
		assert.Equal(t, string(transcription.Status), dto.Status)
		assert.Equal(t, transcription.Duration, dto.Duration)
		assert.Equal(t, transcription.CreatedAt, dto.CreatedAt)
	})

	t.Run("Convert nil transcription", func(t *testing.T) {
		dto := mapper.ToDTO(nil)
		assert.Nil(t, dto)
	})
}

func TestTranscriptionMapper_ToDTOs(t *testing.T) {
	mapper := NewTranscriptionMapper()

	transcriptions := []*entities.Transcription{
		{
			ID:        uuid.New(),
			UserID:    uuid.New(),
			Text:      "First transcription",
			Status:    entities.StatusCompleted,
			Duration:  10.5,
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			UserID:    uuid.New(),
			Text:      "Second transcription",
			Status:    entities.StatusProcessing,
			Duration:  0,
			CreatedAt: time.Now(),
		},
	}

	dtos := mapper.ToDTOs(transcriptions)

	assert.Len(t, dtos, 2)
	assert.Equal(t, transcriptions[0].Text, dtos[0].Text)
	assert.Equal(t, transcriptions[1].Text, dtos[1].Text)
}
