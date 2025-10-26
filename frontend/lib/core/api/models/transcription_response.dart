class TranscriptionResponse {
  final String id;
  final String userId;
  final String text;
  final String status;
  final DateTime createdAt;
  final double duration;

  TranscriptionResponse({
    required this.id,
    required this.userId,
    required this.text,
    required this.status,
    required this.createdAt,
    required this.duration,
  });

  factory TranscriptionResponse.fromJson(Map<String, dynamic> json) {
    return TranscriptionResponse(
      id: json['id']?.toString() ?? '',
      userId: json['user_id']?.toString() ?? '',
      text: json['text']?.toString() ?? '',
      status: json['status']?.toString() ?? 'processing',
      createdAt: json['created_at'] != null 
          ? DateTime.parse(json['created_at'] as String)
          : DateTime.now(),
      duration: (json['duration'] as num?)?.toDouble() ?? 0.0,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'text': text,
      'status': status,
      'created_at': createdAt.toIso8601String(),
      'duration': duration,
    };
  }
}

