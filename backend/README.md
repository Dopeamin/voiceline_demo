# Voiceline Backend

Go backend with Gin framework for voice transcription service.

## Setup

```bash
# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env with your configuration
# Add your OPENAI_API_KEY

# Run the server
go run cmd/server/main.go
```

## Architecture

**Clean Architecture** with clear separation of concerns:

### Domain Layer (`internal/domain/`)
- Entities with business logic

### Application Layer (`internal/application/`)
- Services
- Repositories interfaces

### Infrastructure Layer (`internal/infrastructure/`)
- OpenAI integration
- In-memory repositories

### Interface Layer (`internal/interface/`)
- HTTP handlers
- Middleware
- DTOs and mappers

## API Endpoints

### Health
- `GET /api/v1/health` - Health check

### Auth
- `POST /api/v1/auth/register` - Register user
- `POST /api/v1/auth/login` - Login user

### Transcriptions (Protected)
- `POST /api/v1/transcriptions` - Transcribe audio
- `GET /api/v1/transcriptions` - Get all transcriptions
- `GET /api/v1/transcriptions/:id` - Get transcription by ID

## Testing

All tests are organized in the `tests/` directory:

```bash
# Run all tests
go test ./tests/...

# Run only unit tests
go test ./tests/unit/...

# Run only integration tests
go test ./tests/integration/...

# Run specific test suite
go test -v ./tests/unit/entities/
go test -v ./tests/unit/mappers/

**Test Structure:**
- `tests/unit/entities/` - Entity business logic tests
- `tests/unit/mappers/` - DTO mapper tests
- `tests/integration/` - API endpoint tests

For more information, see the main [README](../README.md).

