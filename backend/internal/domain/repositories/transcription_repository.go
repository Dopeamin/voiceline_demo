package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/domain/entities"
)

type TranscriptionRepository interface {
	Create(ctx context.Context, transcription *entities.Transcription) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Transcription, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Transcription, error)
	Update(ctx context.Context, transcription *entities.Transcription) error
	Delete(ctx context.Context, id uuid.UUID) error
}
