import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:voiceline/core/api/api_client.dart';
import 'package:voiceline/core/api/services/transcription_service.dart';
import 'package:voiceline/core/api/models/transcription_response.dart';
import 'package:voiceline/core/context/auth_context.dart';
import 'package:voiceline/features/recording/screens/transcription_detail_screen.dart';

class TranscriptionList extends StatefulWidget {
  const TranscriptionList({super.key});

  @override
  State<TranscriptionList> createState() => _TranscriptionListState();
}

class _TranscriptionListState extends State<TranscriptionList> {
  late final ApiClient _apiClient;
  late final TranscriptionService _transcriptionService;
  List<TranscriptionResponse> _transcriptions = [];
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    _apiClient = ApiClient();
    _transcriptionService = TranscriptionService(_apiClient);
    _loadTranscriptions();
  }

  Future<void> _loadTranscriptions() async {
    try {
      final token = context.read<AuthContext>().token;
      if (token == null) return;

      _apiClient.setAuthToken(token);

      final transcriptions = await _transcriptionService.getTranscriptions();
      
      if (mounted) {
        setState(() {
          _transcriptions = transcriptions;
          _isLoading = false;
        });
      }
    } catch (e) {
      if (mounted) {
        setState(() => _isLoading = false);
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Failed to load transcriptions: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Center(child: CircularProgressIndicator());
    }

    if (_transcriptions.isEmpty) {
      return const Center(
        child: Text('No transcriptions yet'),
      );
    }

    return RefreshIndicator(
      onRefresh: _loadTranscriptions,
      child: ListView.builder(
        padding: const EdgeInsets.all(16),
        itemCount: _transcriptions.length,
        itemBuilder: (context, index) {
          final transcription = _transcriptions[index];
          return Card(
            margin: const EdgeInsets.only(bottom: 16),
            child: ListTile(
              onTap: () {
                Navigator.of(context).push(
                  MaterialPageRoute(
                    builder: (context) => TranscriptionDetailScreen(
                      transcription: transcription,
                    ),
                    fullscreenDialog: true,
                  ),
                );
              },
              title: Text(
                transcription.text.isEmpty ? 'No transcription' : transcription.text,
                maxLines: 2,
                overflow: TextOverflow.ellipsis,
              ),
              subtitle: Text(
                'Status: ${transcription.status}',
                style: TextStyle(
                  color: transcription.status == 'completed'
                      ? Colors.green
                      : Colors.orange,
                ),
              ),
              trailing: const Icon(Icons.chevron_right),
            ),
          );
        },
      ),
    );
  }
}

