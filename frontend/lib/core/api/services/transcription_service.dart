import 'package:dio/dio.dart';
import '../api_client.dart';
import '../models/transcription_response.dart';

class TranscriptionService {
  final ApiClient _client;

  TranscriptionService(this._client);

  Future<TranscriptionResponse> transcribeAudio({
    required String audioPath,
  }) async {
    final formData = FormData.fromMap({
      'audio': await MultipartFile.fromFile(
        audioPath,
        filename: audioPath.split('/').last,
      ),
    });

    final response = await _client.postMultipart(
      '/transcriptions',
      formData: formData,
    );

    return TranscriptionResponse.fromJson(
      response.data as Map<String, dynamic>,
    );
  }

  Future<List<TranscriptionResponse>> getTranscriptions() async {
    final response = await _client.get('/transcriptions');

    final List<dynamic> data = response.data as List<dynamic>;
    return data
        .map((json) =>
            TranscriptionResponse.fromJson(json as Map<String, dynamic>))
        .toList();
  }

  Future<TranscriptionResponse> getTranscription(String id) async {
    final response = await _client.get('/transcriptions/$id');

    return TranscriptionResponse.fromJson(
      response.data as Map<String, dynamic>,
    );
  }
}

