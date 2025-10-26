package openai

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	ErrEmptyAPIKey          = errors.New("OpenAI API key is empty")
	ErrTranscriptionFailed  = errors.New("transcription failed")
	ErrServiceNotConfigured = errors.New("OpenAI service is not configured. Please set OPENAI_API_KEY environment variable")
)

type TranscriptionService struct {
	client *openai.Client
}

// NewTranscriptionService creates a new TranscriptionService
func NewTranscriptionService(apiKey string) (*TranscriptionService, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	client := openai.NewClient(apiKey)

	return &TranscriptionService{
		client: client,
	}, nil
}

func (s *TranscriptionService) TranscribeAudio(ctx context.Context, audio io.Reader) (string, float64, error) {

	if s.client == nil {
		return "", 0, ErrServiceNotConfigured
	}

	// OpenAI SDK requires a file path
	tmpFile, err := os.CreateTemp("", "audio-*.m4a")
	if err != nil {
		return "", 0, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, audio)
	if err != nil {
		return "", 0, err
	}

	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return "", 0, err
	}

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: tmpFile.Name(),
		Format:   openai.AudioResponseFormatVerboseJSON,
		Language: "en", // Specify English for better accuracy
	}

	resp, err := s.client.CreateTranscription(ctx, req)
	if err != nil {
		return "", 0, err
	}

	if resp.Text == "" {
		return "", 0, ErrTranscriptionFailed
	}

	// Use duration from OpenAI response
	duration := float64(resp.Duration)

	return resp.Text, duration, nil
}
