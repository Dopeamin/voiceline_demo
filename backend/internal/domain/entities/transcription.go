package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TranscriptionStatus string

const (
	StatusProcessing TranscriptionStatus = "processing"
	StatusCompleted  TranscriptionStatus = "completed"
	StatusFailed     TranscriptionStatus = "failed"
)

var (
	ErrInvalidTranscriptionStatus = errors.New("invalid transcription status")
	ErrEmptyText                  = errors.New("transcription text cannot be empty")
)

type Transcription struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Text      string
	Status    TranscriptionStatus
	Duration  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTranscription(userID uuid.UUID) *Transcription {
	now := time.Now()
	return &Transcription{
		ID:        uuid.New(),
		UserID:    userID,
		Status:    StatusProcessing,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Transcription) Complete(text string, duration float64) error {
	if text == "" {
		return ErrEmptyText
	}

	t.Text = text
	t.Duration = duration
	t.Status = StatusCompleted
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transcription) Fail() {
	t.Status = StatusFailed
	t.UpdatedAt = time.Now()
}

func (t *Transcription) IsCompleted() bool {
	return t.Status == StatusCompleted
}

func (t *Transcription) IsFailed() bool {
	return t.Status == StatusFailed
}

func (t *Transcription) IsProcessing() bool {
	return t.Status == StatusProcessing
}

func (t *Transcription) BelongsToUser(userID uuid.UUID) bool {
	return t.UserID == userID
}
