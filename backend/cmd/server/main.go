package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/infrastructure/openai"
	"github.com/voiceline/backend/internal/infrastructure/persistence"
	httpInterface "github.com/voiceline/backend/internal/interface/http"
)

func main() {
	// Load .env file if it exists (ignore error if file doesn't exist)
	_ = godotenv.Load()

	// Load configuration from environment
	port := getEnv("PORT", "8080")
	jwtSecret := getEnv("JWT_SECRET", "your-super-secret-jwt-key")
	openAIKey := getEnv("OPENAI_API_KEY", "")

	if openAIKey == "" {
		log.Println("WARNING: OPENAI_API_KEY is not set. Transcription will not work.")
	}

	// Initialize repositories
	userRepo := persistence.NewMemoryUserRepository()
	transcriptionRepo := persistence.NewMemoryTranscriptionRepository()

	// Initialize OpenAI service
	openAIService, err := openai.NewTranscriptionService(openAIKey)
	if err != nil {
		log.Fatalf("Failed to initialize OpenAI service: %v", err)
	}

	if openAIKey == "" {
		log.Println("WARNING: OPENAI_API_KEY is not set. Transcription endpoints will return an error.")
	}

	// Initialize services
	authService := services.NewAuthService(userRepo, jwtSecret)
	transcriptionService := services.NewTranscriptionService(transcriptionRepo, openAIService)

	// Initialize HTTP router
	router := httpInterface.NewRouter(authService, transcriptionService)
	engine := router.Setup()

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting on %s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
