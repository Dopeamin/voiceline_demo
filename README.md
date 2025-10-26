# Voiceline 🎤

A mobile voice recording and transcription app built with **Flutter** and **Go**.

Record audio on your phone → Send to backend → Get AI transcription → View results.

> **📱 Platform**: Tested on **iOS 13.0+** only. Android support not yet verified.

https://github.com/user-attachments/assets/dd5bac75-c66e-41cd-87e9-5a9e1a18206c

---

## 🏗️ Architecture

### Backend (Go + Gin)

Clean architecture with clear separation of concerns:

```
backend/
├── cmd/server/          # Application entry point
├── internal/
│   ├── domain/          # Business logic & entities
│   │   ├── entities/    # User, Transcription
│   │   └── repositories/  # Repository interfaces
│   ├── application/     # Business services
│   │   └── services/    # AuthService, TranscriptionService
│   ├── infrastructure/  # External services
│   │   ├── openai/      # OpenAI Whisper integration
│   │   └── persistence/ # In-memory repositories
│   └── interface/       # HTTP layer
│       ├── dto/         # Data transfer objects
│       ├── mappers/     # Entity ↔ DTO conversion
│       └── http/        # Handlers, middleware, routes
└── tests/
    ├── unit/            # Entity & mapper tests
    └── integration/     # API endpoint tests
```

**Key Features:**
- JWT authentication
- Clean architecture (domain → application → infrastructure → interface)
- OpenAI Whisper for transcription
- In-memory storage (easy to swap for database)
- Comprehensive tests

### Frontend (Flutter)

Feature-based structure with global state management:

```
frontend/lib/
├── core/
│   ├── api/
│   │   ├── models/      # User, AuthResponse, TranscriptionResponse
│   │   ├── services/    # AuthService, TranscriptionService
│   │   └── api_client.dart  # Dio HTTP client
│   ├── context/         # AuthContext (global auth state)
│   ├── navigation/      # App routing logic
│   └── shared/          # Reusable UI components
└── features/
    ├── auth/            # Login & Registration
    └── recording/       # Audio recording & transcriptions
```

**Key Features:**
- Provider for state management
- Secure token storage
- Auto token validation on startup
- Manual API client (simple, no code generation)
- Clean separation of features

---

## 🚀 Quick Start

### Prerequisites

- **Go** 1.21+
- **Flutter** 3.0+
- **OpenAI API Key** (optional, uses mock mode without it)
- **iOS Device/Simulator** (iOS 13.0+)

> ⚠️ **Note**: Due to time constraints, this app has been tested **only on iOS** (latest versions). While Flutter is cross-platform, Android compatibility has not been verified. We recommend running the app on **iOS 13.0 or later** for the best experience.

### Backend Setup

```bash
cd backend

# Create .env file (copy from example or create new)
cat > .env << EOF
PORT=8080
JWT_SECRET=your-super-secret-jwt-key-change-this
OPENAI_API_KEY=sk-your-key-here
EOF

# Run server
go run cmd/server/main.go
```

Server runs on `http://localhost:8080`

**Without OpenAI Key:** Leave `OPENAI_API_KEY` empty or omit it - mock transcriptions will be returned for testing.

### Frontend Setup

```bash
cd frontend

# Install dependencies
flutter pub get

# Run on iOS simulator (recommended - tested platform)
flutter run -d ios

# Or run on a specific iOS device
flutter devices  # List available devices
flutter run -d <device-id>
```

**Note:** Update the base URL in `lib/core/api/api_client.dart`:
- iOS Simulator: `http://localhost:8080/api/v1`
- Android Emulator: `http://10.0.2.2:8080/api/v1`
- Physical Device: `http://YOUR_IP:8080/api/v1`

### Test It Out

1. Register a new account
2. Login
3. Record audio
4. View transcription (mock or real)
5. See your transcription history

---

## 🧪 Testing

```bash
# Backend - Run all tests
cd backend
go test ./tests/...

# Backend - Unit tests only
go test ./tests/unit/...

# Backend - Integration tests only
go test ./tests/integration/...
```

---

## 🔐 Authentication Flow

```
User → Register/Login → JWT Token → Secure Storage
                                          ↓
App Restart → Load Token → Validate with Backend
                                ↓
                          Valid? Stay logged in : Logout
```

- Tokens validated automatically on app startup
- Invalid/expired tokens cleared automatically
- Seamless user experience

---

## 📡 API Endpoints

### Public

- `POST /api/v1/auth/register` - Create account
- `POST /api/v1/auth/login` - Get JWT token
- `GET /api/v1/health` - Health check

### Protected (requires JWT)

- `POST /api/v1/transcriptions` - Upload audio for transcription
- `GET /api/v1/transcriptions` - Get user's transcriptions
- `GET /api/v1/transcriptions/:id` - Get specific transcription

---

## 🎯 Next Steps / Potential Improvements

- [ ] **Batch Upload**: Upload multiple audio files
- [ ] **Voice Commands**: Control app with voice
- [ ] **Speaker Identification**: Detect different speakers
- [ ] **Summarization**: AI-powered summaries of long transcriptions
- [ ] **Translation**: Translate transcriptions to other languages
- [ ] **Collaboration**: Share transcriptions with other users
- [ ] **Analytics**: Usage statistics and insights

---

## 📝 Configuration

### Backend `.env`

Create a `.env` file in the `backend/` directory:

```bash
PORT=8080
JWT_SECRET=your-super-secret-jwt-key-change-this
OPENAI_API_KEY=sk-your-openai-key-here
```

**Note**: The `.env` file is automatically loaded by the backend. Never commit this file to version control.

### Flutter API Client

Edit `frontend/lib/core/api/api_client.dart`:

```dart
ApiClient({String baseUrl = 'http://localhost:8080/api/v1'})
```
