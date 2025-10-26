package services

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/domain/entities"
	"github.com/voiceline/backend/internal/domain/repositories"
)

var (
	ErrTranscriptionNotFound = errors.New("transcription not found")
	ErrUnauthorizedAccess    = errors.New("unauthorized access to transcription")
)

type ITranscriptionService interface {
	TranscribeAudio(ctx context.Context, audio io.Reader) (string, float64, error)
}

type TranscriptionService struct {
	transcriptionRepo repositories.TranscriptionRepository
	transcriptionSvc  ITranscriptionService
}

func NewTranscriptionService(
	transcriptionRepo repositories.TranscriptionRepository,
	transcriptionSvc ITranscriptionService,
) *TranscriptionService {
	return &TranscriptionService{
		transcriptionRepo: transcriptionRepo,
		transcriptionSvc:  transcriptionSvc,
	}
}

type TranscribeAudioInput struct {
	UserID uuid.UUID
	Audio  io.Reader
}

func (s *TranscriptionService) Transcribe(ctx context.Context, input TranscribeAudioInput) (*entities.Transcription, error) {
	transcription := entities.NewTranscription(input.UserID)

	if err := s.transcriptionRepo.Create(ctx, transcription); err != nil {
		return nil, err
	}

	text, duration, err := s.transcriptionSvc.TranscribeAudio(ctx, input.Audio)
	if err != nil {
		transcription.Fail()
		_ = s.transcriptionRepo.Update(ctx, transcription)
		return nil, err
	}

	if err := transcription.Complete(text, duration); err != nil {
		transcription.Fail()
		_ = s.transcriptionRepo.Update(ctx, transcription)
		return nil, err
	}

	if err := s.transcriptionRepo.Update(ctx, transcription); err != nil {
		return nil, err
	}

	return transcription, nil
}

func (s *TranscriptionService) GetTranscription(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*entities.Transcription, error) {
	transcription, err := s.transcriptionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrTranscriptionNotFound
	}

	if !transcription.BelongsToUser(userID) {
		return nil, ErrUnauthorizedAccess
	}

	return transcription, nil
}

func (s *TranscriptionService) GetUserTranscriptions(ctx context.Context, userID uuid.UUID) ([]*entities.Transcription, error) {
	transcriptions, err := s.transcriptionRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return transcriptions, nil
}
