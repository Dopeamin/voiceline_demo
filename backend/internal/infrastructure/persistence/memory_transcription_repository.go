package persistence

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/domain/entities"
)

var (
	ErrTranscriptionNotFound = errors.New("transcription not found")
)

type MemoryTranscriptionRepository struct {
	transcriptions map[uuid.UUID]*entities.Transcription
	userIndex      map[uuid.UUID][]uuid.UUID
	mu             sync.RWMutex
}

func NewMemoryTranscriptionRepository() *MemoryTranscriptionRepository {
	return &MemoryTranscriptionRepository{
		transcriptions: make(map[uuid.UUID]*entities.Transcription),
		userIndex:      make(map[uuid.UUID][]uuid.UUID),
	}
}

func (r *MemoryTranscriptionRepository) Create(ctx context.Context, transcription *entities.Transcription) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.transcriptions[transcription.ID] = transcription
	r.userIndex[transcription.UserID] = append(r.userIndex[transcription.UserID], transcription.ID)
	return nil
}

func (r *MemoryTranscriptionRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Transcription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transcription, exists := r.transcriptions[id]
	if !exists {
		return nil, ErrTranscriptionNotFound
	}

	return transcription, nil
}

func (r *MemoryTranscriptionRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Transcription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ids, exists := r.userIndex[userID]
	if !exists {
		return []*entities.Transcription{}, nil
	}

	transcriptions := make([]*entities.Transcription, 0, len(ids))
	for _, id := range ids {
		if transcription, exists := r.transcriptions[id]; exists {
			transcriptions = append(transcriptions, transcription)
		}
	}

	return transcriptions, nil
}

func (r *MemoryTranscriptionRepository) Update(ctx context.Context, transcription *entities.Transcription) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.transcriptions[transcription.ID]; !exists {
		return ErrTranscriptionNotFound
	}

	r.transcriptions[transcription.ID] = transcription
	return nil
}

func (r *MemoryTranscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	transcription, exists := r.transcriptions[id]
	if !exists {
		return ErrTranscriptionNotFound
	}

	userIDs := r.userIndex[transcription.UserID]
	for i, tID := range userIDs {
		if tID == id {
			r.userIndex[transcription.UserID] = append(userIDs[:i], userIDs[i+1:]...)
			break
		}
	}

	delete(r.transcriptions, id)
	return nil
}
