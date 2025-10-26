package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/voiceline/backend/internal/application/services"
	"github.com/voiceline/backend/internal/interface/http/handlers"
	"github.com/voiceline/backend/internal/interface/http/middleware"
	"github.com/voiceline/backend/internal/interface/mappers"
)

// Router sets up the HTTP router
type Router struct {
	engine               *gin.Engine
	authService          *services.AuthService
	transcriptionService *services.TranscriptionService
}

// NewRouter creates a new HTTP router
func NewRouter(
	authService *services.AuthService,
	transcriptionService *services.TranscriptionService,
) *Router {
	return &Router{
		engine:               gin.Default(),
		authService:          authService,
		transcriptionService: transcriptionService,
	}
}

// Setup configures all routes
func (r *Router) Setup() *gin.Engine {
	// Configure CORS
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize mappers
	userMapper := mappers.NewUserMapper()
	transcriptionMapper := mappers.NewTranscriptionMapper()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler("1.0.0")
	authHandler := handlers.NewAuthHandler(r.authService, userMapper)
	transcriptionHandler := handlers.NewTranscriptionHandler(r.transcriptionService, transcriptionMapper)

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", healthHandler.HealthCheck)

		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Transcription routes (protected)
		transcriptions := v1.Group("/transcriptions")
		transcriptions.Use(middleware.AuthMiddleware(r.authService))
		{
			transcriptions.POST("", transcriptionHandler.TranscribeAudio)
			transcriptions.GET("", transcriptionHandler.GetTranscriptions)
			transcriptions.GET("/:id", transcriptionHandler.GetTranscription)
		}
	}

	return r.engine
}
