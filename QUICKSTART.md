# Quick Start Guide

Get Voiceline up and running in 5 minutes!

## Prerequisites

- Go 1.21+
- Flutter 3.0+
- OpenAI API Key

## Backend Quick Start

```bash
# 1. Navigate to backend
cd backend

# 2. Copy environment file
cp .env.example .env

# 3. Edit .env and add your OpenAI API key
# OPENAI_API_KEY=your-key-here

# 4. Install dependencies
go mod download

# 5. Run the server
go run cmd/server/main.go
```

Backend will be running at `http://localhost:8080`

## Frontend Quick Start

```bash
cd frontend

# Install dependencies
flutter pub get

# Run app
flutter run
```

**API Endpoint:** Edit `lib/core/api/api_client.dart` if needed:
- iOS Simulator: `localhost:8080`
- Android Emulator: `10.0.2.2:8080`
- Physical Device: `YOUR_IP:8080`

## Test It Out

1. **Register** a new account in the app
2. **Login** with your credentials
3. **Tap the microphone button** to record
4. **Speak** something clearly
5. **Tap again** to stop and see the transcription!

## Troubleshooting

**Can't connect from app to backend?**
- Android emulator: Use `http://10.0.2.2:8080/api/v1`
- Check firewall settings
- Ensure backend is running

**Audio recording not working?**
- Grant microphone permissions
- Test on a physical device (emulators may not support audio)

**Transcription fails?**
- Verify OpenAI API key is correct
- Check your OpenAI account has credits
- Ensure audio file format is supported

For more detailed instructions, see [README.md](./README.md)

