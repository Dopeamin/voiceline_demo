package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/voiceline/backend/internal/domain/entities"
)

var (
	NewTranscription = entities.NewTranscription
	StatusProcessing = entities.StatusProcessing
	StatusCompleted  = entities.StatusCompleted
	StatusFailed     = entities.StatusFailed
	ErrEmptyText     = entities.ErrEmptyText
)

func TestNewTranscription(t *testing.T) {
	userID := uuid.New()
	transcription := NewTranscription(userID)

	assert.NotNil(t, transcription)
	assert.NotEqual(t, uuid.Nil, transcription.ID)
	assert.Equal(t, userID, transcription.UserID)
	assert.Equal(t, StatusProcessing, transcription.Status)
	assert.Empty(t, transcription.Text)
	assert.Equal(t, 0.0, transcription.Duration)
}

func TestTranscription_Complete(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		duration    float64
		expectError bool
	}{
		{
			name:        "Valid completion",
			text:        "This is a transcribed text",
			duration:    10.5,
			expectError: false,
		},
		{
			name:        "Empty text",
			text:        "",
			duration:    10.5,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trans := NewTranscription(uuid.New())
			err := trans.Complete(tt.text, tt.duration)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, ErrEmptyText, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.text, trans.Text)
				assert.Equal(t, tt.duration, trans.Duration)
				assert.Equal(t, StatusCompleted, trans.Status)
			}
		})
	}
}

func TestTranscription_Fail(t *testing.T) {
	transcription := NewTranscription(uuid.New())
	transcription.Fail()

	assert.Equal(t, StatusFailed, transcription.Status)
}

func TestTranscription_StatusChecks(t *testing.T) {
	userID := uuid.New()

	t.Run("Processing status", func(t *testing.T) {
		trans := NewTranscription(userID)
		assert.True(t, trans.IsProcessing())
		assert.False(t, trans.IsCompleted())
		assert.False(t, trans.IsFailed())
	})

	t.Run("Completed status", func(t *testing.T) {
		trans := NewTranscription(userID)
		_ = trans.Complete("Test text", 10.0)
		assert.False(t, trans.IsProcessing())
		assert.True(t, trans.IsCompleted())
		assert.False(t, trans.IsFailed())
	})

	t.Run("Failed status", func(t *testing.T) {
		trans := NewTranscription(userID)
		trans.Fail()
		assert.False(t, trans.IsProcessing())
		assert.False(t, trans.IsCompleted())
		assert.True(t, trans.IsFailed())
	})
}

func TestTranscription_BelongsToUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()

	transcription := NewTranscription(userID)

	assert.True(t, transcription.BelongsToUser(userID))
	assert.False(t, transcription.BelongsToUser(otherUserID))
}
