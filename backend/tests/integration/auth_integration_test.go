package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/infrastructure/persistence"
	httpInterface "github.com/voiceline/backend/internal/interface/http"
)

func setupTestServer() *httptest.Server {
	userRepo := persistence.NewMemoryUserRepository()
	authService := services.NewAuthService(userRepo, "test-secret")

	transcriptionRepo := persistence.NewMemoryTranscriptionRepository()
	transcriptionService := services.NewTranscriptionService(transcriptionRepo, nil)

	router := httpInterface.NewRouter(authService, transcriptionService)
	engine := router.Setup()

	return httptest.NewServer(engine)
}

func TestAuthIntegration_Register(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	tests := []struct {
		name           string
		payload        map[string]string
		expectedStatus int
	}{
		{
			name: "Successful registration",
			payload: map[string]string{
				"email":    "test@example.com",
				"password": "password123",
				"name":     "Test User",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Invalid email",
			payload: map[string]string{
				"email":    "invalid-email",
				"password": "password123",
				"name":     "Test User",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Short password",
			payload: map[string]string{
				"email":    "test2@example.com",
				"password": "short",
				"name":     "Test User",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			resp, err := http.Post(server.URL+"/api/v1/auth/register", "application/json", bytes.NewBuffer(body))

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus == http.StatusCreated {
				var result map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&result)
				assert.NotEmpty(t, result["token"])
				assert.NotNil(t, result["user"])
			}
		})
	}
}

func TestAuthIntegration_Login(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	registerPayload := map[string]string{
		"email":    "login@example.com",
		"password": "password123",
		"name":     "Login User",
	}
	body, _ := json.Marshal(registerPayload)
	http.Post(server.URL+"/api/v1/auth/register", "application/json", bytes.NewBuffer(body))

	tests := []struct {
		name           string
		payload        map[string]string
		expectedStatus int
	}{
		{
			name: "Successful login",
			payload: map[string]string{
				"email":    "login@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Wrong password",
			payload: map[string]string{
				"email":    "login@example.com",
				"password": "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Non-existent user",
			payload: map[string]string{
				"email":    "nonexistent@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			resp, err := http.Post(server.URL+"/api/v1/auth/login", "application/json", bytes.NewBuffer(body))

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus == http.StatusOK {
				var result map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&result)
				assert.NotEmpty(t, result["token"])
			}
		})
	}
}
