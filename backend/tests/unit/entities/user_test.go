package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voiceline/backend/internal/domain/entities"
)

var (
	NewUser            = entities.NewUser
	ErrInvalidEmail    = entities.ErrInvalidEmail
	ErrInvalidPassword = entities.ErrInvalidPassword
	ErrInvalidName     = entities.ErrInvalidName
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		userName    string
		expectError bool
		errorType   error
	}{
		{
			name:        "Valid user",
			email:       "test@example.com",
			password:    "password123",
			userName:    "Test User",
			expectError: false,
		},
		{
			name:        "Invalid email",
			email:       "invalid-email",
			password:    "password123",
			userName:    "Test User",
			expectError: true,
			errorType:   ErrInvalidEmail,
		},
		{
			name:        "Short password",
			email:       "test@example.com",
			password:    "short",
			userName:    "Test User",
			expectError: true,
			errorType:   ErrInvalidPassword,
		},
		{
			name:        "Empty name",
			email:       "test@example.com",
			password:    "password123",
			userName:    "",
			expectError: true,
			errorType:   ErrInvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.email, tt.password, tt.userName)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorType, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.userName, user.Name)
				assert.NotEmpty(t, user.ID)
				assert.NotEmpty(t, user.PasswordHash)
			}
		})
	}
}

func TestUser_VerifyPassword(t *testing.T) {
	user, err := NewUser("test@example.com", "password123", "Test User")
	assert.NoError(t, err)

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Correct password",
			password: "password123",
			expected: true,
		},
		{
			name:     "Wrong password",
			password: "wrongpassword",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := user.VerifyPassword(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUser_UpdatePassword(t *testing.T) {
	user, err := NewUser("test@example.com", "password123", "Test User")
	assert.NoError(t, err)

	// Update to valid password
	err = user.UpdatePassword("newpassword123")
	assert.NoError(t, err)
	assert.True(t, user.VerifyPassword("newpassword123"))
	assert.False(t, user.VerifyPassword("password123"))

	// Try to update to invalid password
	err = user.UpdatePassword("short")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidPassword, err)
}
