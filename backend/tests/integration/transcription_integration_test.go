package integration

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAuthToken(server *httptest.Server) string {
	registerPayload := map[string]string{
		"email":    "transcription@example.com",
		"password": "password123",
		"name":     "Transcription User",
	}
	body, _ := json.Marshal(registerPayload)
	resp, _ := http.Post(server.URL+"/api/v1/auth/register", "application/json", bytes.NewBuffer(body))

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["token"].(string)
}

func TestTranscriptionIntegration_Unauthorized(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL+"/api/v1/transcriptions", nil)
	resp, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestTranscriptionIntegration_GetTranscriptions(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	token := getAuthToken(server)

	req, _ := http.NewRequest("GET", server.URL+"/api/v1/transcriptions", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var transcriptions []interface{}
	json.NewDecoder(resp.Body).Decode(&transcriptions)
	assert.Empty(t, transcriptions)
}

func TestTranscriptionIntegration_TranscribeAudio_MissingFile(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	token := getAuthToken(server)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("POST", server.URL+"/api/v1/transcriptions", body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
