package mappers

import (
	"github.com/voiceline/backend/internal/domain/entities"
	"github.com/voiceline/backend/internal/interface/dto"
)

// TranscriptionMapper handles mapping between Transcription entity and DTOs
type TranscriptionMapper struct{}

// NewTranscriptionMapper creates a new TranscriptionMapper
func NewTranscriptionMapper() *TranscriptionMapper {
	return &TranscriptionMapper{}
}

// ToDTO converts a Transcription entity to a TranscriptionDTO
func (m *TranscriptionMapper) ToDTO(transcription *entities.Transcription) *dto.TranscriptionDTO {
	if transcription == nil {
		return nil
	}

	return &dto.TranscriptionDTO{
		ID:        transcription.ID.String(),
		UserID:    transcription.UserID.String(),
		Text:      transcription.Text,
		Status:    string(transcription.Status),
		Duration:  transcription.Duration,
		CreatedAt: transcription.CreatedAt,
	}
}

// ToDTOs converts a slice of Transcription entities to TranscriptionDTOs
func (m *TranscriptionMapper) ToDTOs(transcriptions []*entities.Transcription) []*dto.TranscriptionDTO {
	dtos := make([]*dto.TranscriptionDTO, len(transcriptions))
	for i, transcription := range transcriptions {
		dtos[i] = m.ToDTO(transcription)
	}
	return dtos
}
