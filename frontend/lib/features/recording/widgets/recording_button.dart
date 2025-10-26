import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:provider/provider.dart';
import 'package:record/record.dart';
import 'package:voiceline/core/api/api_client.dart';
import 'package:voiceline/core/api/services/transcription_service.dart';
import 'package:voiceline/core/context/auth_context.dart';
import 'package:voiceline/features/recording/screens/transcription_detail_screen.dart';

class RecordingButton extends StatefulWidget {
  const RecordingButton({super.key});

  @override
  State<RecordingButton> createState() => _RecordingButtonState();
}

class _RecordingButtonState extends State<RecordingButton> {
  final AudioRecorder _recorder = AudioRecorder();
  late final ApiClient _apiClient;
  late final TranscriptionService _transcriptionService;
  
  bool _isRecording = false;
  bool _isProcessing = false;

  @override
  void initState() {
    super.initState();
    _apiClient = ApiClient();
    _transcriptionService = TranscriptionService(_apiClient);
  }

  @override
  void dispose() {
    _recorder.dispose();
    super.dispose();
  }

  Future<void> _requestPermission() async {
    await Permission.microphone.request();
  }

  Future<void> _startRecording() async {
    await _requestPermission();

    if (await _recorder.hasPermission()) {
      final dir = await getTemporaryDirectory();
      final path = '${dir.path}/audio_${DateTime.now().millisecondsSinceEpoch}.m4a';

      await _recorder.start(
        const RecordConfig(encoder: AudioEncoder.aacLc),
        path: path,
      );

      setState(() {
        _isRecording = true;
      });
    }
  }

  Future<void> _stopRecording() async {
    final path = await _recorder.stop();

    setState(() {
      _isRecording = false;
    });

    if (path != null) {
      await _transcribeAudio(path);
    }
  }

  Future<void> _transcribeAudio(String path) async {
    setState(() => _isProcessing = true);

    try {
      final token = context.read<AuthContext>().token;
      if (token == null) return;

      _apiClient.setAuthToken(token);

      final response = await _transcriptionService.transcribeAudio(
        audioPath: path,
      );

      if (mounted) {
        setState(() => _isProcessing = false);
        Navigator.of(context).push(
          MaterialPageRoute(
            builder: (context) => TranscriptionDetailScreen(
              transcription: response,
            ),
            fullscreenDialog: true,
          ),
        );
      }
    } catch (e) {
      if (mounted) {
        setState(() => _isProcessing = false);
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Transcription failed: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isProcessing) {
      return const Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          CircularProgressIndicator(),
          SizedBox(height: 16),
          Text('Transcribing...'),
        ],
      );
    }

    return GestureDetector(
      onTap: _isRecording ? _stopRecording : _startRecording,
      child: Container(
        width: 80,
        height: 80,
        decoration: BoxDecoration(
          color: _isRecording ? Colors.red : Colors.blue,
          shape: BoxShape.circle,
        ),
        child: Icon(
          _isRecording ? Icons.stop : Icons.mic,
          color: Colors.white,
          size: 40,
        ),
      ),
    );
  }
}

