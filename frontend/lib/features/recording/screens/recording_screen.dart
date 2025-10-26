import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:voiceline/core/context/auth_context.dart';
import 'package:voiceline/features/recording/widgets/recording_button.dart';
import 'package:voiceline/features/recording/widgets/transcription_list.dart';

class RecordingScreen extends StatefulWidget {
  const RecordingScreen({super.key});

  @override
  State<RecordingScreen> createState() => _RecordingScreenState();
}

class _RecordingScreenState extends State<RecordingScreen> {
  bool _showHistory = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Voiceline'),
        actions: [
          IconButton(
            icon: Icon(_showHistory ? Icons.mic : Icons.history),
            onPressed: () {
              setState(() {
                _showHistory = !_showHistory;
              });
            },
          ),
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: () {
              context.read<AuthContext>().logout();
            },
          ),
        ],
      ),
      body: AnimatedSwitcher(
        duration: const Duration(milliseconds: 300),
        child: _showHistory
            ? const TranscriptionList(key: ValueKey('history'))
            : const RecordingView(key: ValueKey('recording')),
      ),
    );
  }
}

class RecordingView extends StatelessWidget {
  const RecordingView({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Padding(
        padding: EdgeInsets.all(24.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(
              Icons.mic,
              size: 80,
              color: Colors.blue,
            ),
            SizedBox(height: 24),
            Text(
              'Tap to record your voice',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.w600,
              ),
            ),
            SizedBox(height: 48),
            RecordingButton(),
          ],
        ),
      ),
    );
  }
}

