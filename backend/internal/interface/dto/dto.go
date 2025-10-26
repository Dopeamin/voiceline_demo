package dto

import "time"

// RegisterRequestDTO represents registration request
type RegisterRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequestDTO represents login request
type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserDTO represents user data transfer object
type UserDTO struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// AuthResponseDTO represents authentication response
type AuthResponseDTO struct {
	Token string   `json:"token"`
	User  *UserDTO `json:"user"`
}

// TranscriptionDTO represents transcription data transfer object
type TranscriptionDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	Duration  float64   `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
}

// ErrorDTO represents error response
type ErrorDTO struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

// HealthResponseDTO represents health check response
type HealthResponseDTO struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
