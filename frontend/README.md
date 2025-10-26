# Voiceline Frontend

Flutter mobile app for voice recording and transcription.

## Setup

```bash
# Install dependencies
flutter pub get

# Run the app
flutter run
```

## API Structure

The app uses a clean API architecture:

- **Models** (`lib/core/api/models/`): Data models for API requests/responses
  - `User`, `AuthResponse`, `TranscriptionResponse`
- **Services** (`lib/core/api/services/`): Service classes for each API domain
  - `AuthService`: Handles registration and login
  - `TranscriptionService`: Handles audio transcription
- **ApiClient** (`lib/core/api/api_client.dart`): Dio-based HTTP client with auth token management

## Configuration

Update the API base URL in `lib/core/api/api_client.dart`:

- **Android Emulator**: `http://10.0.2.2:8080/api/v1`
- **iOS Simulator**: `http://localhost:8080/api/v1`
- **Physical Device**: `http://YOUR_COMPUTER_IP:8080/api/v1`

## Features

- User authentication (register/login)
- Audio recording with real-time feedback
- Transcription via OpenAI Whisper
- Transcription history
- Clean, modern UI

## Architecture

```
lib/
├── core/
│   ├── api/          # API client
│   ├── context/      # Global state (Provider)
│   ├── navigation/   # App routing
│   └── shared/       # Reusable widgets
└── features/
    ├── auth/         # Authentication
    └── recording/    # Recording & transcription
```

## Permissions

### Android
Permissions are declared in `android/app/src/main/AndroidManifest.xml`:
- Microphone
- Internet
- Storage

### iOS
Permissions are declared in `ios/Runner/Info.plist`:
- Microphone usage description

## Dependencies

Key packages:
- `provider` - State management
- `record` - Audio recording
- `http` - API communication
- `flutter_secure_storage` - Secure token storage
- `permission_handler` - Permission handling

## Development

```bash
# Run with hot reload
flutter run

# Run tests (if available)
flutter test

# Build for release
flutter build apk         # Android
flutter build ios         # iOS
```

For more information, see the main [README](../README.md).

