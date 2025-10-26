import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:voiceline/core/api/models/transcription_response.dart';
import 'package:intl/intl.dart';

class TranscriptionDetailScreen extends StatelessWidget {
  final TranscriptionResponse transcription;

  const TranscriptionDetailScreen({
    super.key,
    required this.transcription,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.transparent,
      body: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
            colors: [
              Colors.blue.shade50,
              Colors.purple.shade50,
            ],
          ),
        ),
        child: SafeArea(
          child: Column(
            children: [
              _buildHeader(context),
              Expanded(
                child: SingleChildScrollView(
                  padding: const EdgeInsets.all(24),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      _buildStatusCard(),
                      const SizedBox(height: 16),
                      _buildTranscriptionCard(context),
                      const SizedBox(height: 16),
                      _buildMetadataCard(),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildHeader(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16),
      child: Row(
        children: [
          IconButton(
            icon: const Icon(Icons.close, size: 28),
            onPressed: () => Navigator.of(context).pop(),
          ),
          const Spacer(),
          const Text(
            'Transcription Result',
            style: TextStyle(
              fontSize: 20,
              fontWeight: FontWeight.bold,
            ),
          ),
          const Spacer(),
          const SizedBox(width: 48),
        ],
      ),
    );
  }

  Widget _buildStatusCard() {
    final statusInfo = _getStatusInfo();
    return Container(
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: statusInfo.color.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(
          color: statusInfo.color.withValues(alpha: 0.3),
          width: 2,
        ),
      ),
      child: Row(
        children: [
          Container(
            padding: const EdgeInsets.all(12),
            decoration: BoxDecoration(
              color: statusInfo.color.withValues(alpha: 0.2),
              shape: BoxShape.circle,
            ),
            child: Icon(
              statusInfo.icon,
              color: statusInfo.color,
              size: 32,
            ),
          ),
          const SizedBox(width: 16),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  statusInfo.title,
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                    color: statusInfo.color,
                  ),
                ),
                const SizedBox(height: 4),
                Text(
                  statusInfo.subtitle,
                  style: TextStyle(
                    fontSize: 14,
                    color: Colors.grey.shade700,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildTranscriptionCard(BuildContext context) {
    final hasText = transcription.text.isNotEmpty;
    
    return Container(
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(16),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withValues(alpha: 0.05),
            blurRadius: 10,
            offset: const Offset(0, 4),
          ),
        ],
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              const Text(
                'Transcribed Text',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w600,
                  color: Colors.black87,
                ),
              ),
              if (hasText)
                IconButton(
                  icon: const Icon(Icons.copy, size: 20),
                  onPressed: () => _copyToClipboard(context),
                  tooltip: 'Copy to clipboard',
                ),
            ],
          ),
          const SizedBox(height: 16),
          Container(
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: Colors.grey.shade50,
              borderRadius: BorderRadius.circular(12),
              border: Border.all(color: Colors.grey.shade200),
            ),
            child: Text(
              hasText ? transcription.text : 'No transcription available',
              style: TextStyle(
                fontSize: 16,
                height: 1.5,
                color: hasText ? Colors.black87 : Colors.grey.shade500,
                fontStyle: hasText ? FontStyle.normal : FontStyle.italic,
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildMetadataCard() {
    final metadata = [
      _MetadataItem(
        icon: Icons.schedule,
        label: 'Created',
        value: _formatDate(transcription.createdAt),
      ),
      _MetadataItem(
        icon: Icons.timer,
        label: 'Duration',
        value: _formatDuration(transcription.duration),
      ),
      _MetadataItem(
        icon: Icons.fingerprint,
        label: 'ID',
        value: _formatId(transcription.id),
      ),
    ];

    return Container(
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(16),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withValues(alpha: 0.05),
            blurRadius: 10,
            offset: const Offset(0, 4),
          ),
        ],
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'Details',
            style: TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.w600,
              color: Colors.black87,
            ),
          ),
          const SizedBox(height: 16),
          ...metadata.map((item) => _buildMetadataRow(item)),
        ],
      ),
    );
  }

  Widget _buildMetadataRow(_MetadataItem item) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 12),
      child: Row(
        children: [
          Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              color: Colors.blue.shade50,
              borderRadius: BorderRadius.circular(8),
            ),
            child: Icon(
              item.icon,
              size: 20,
              color: Colors.blue.shade700,
            ),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  item.label,
                  style: TextStyle(
                    fontSize: 12,
                    color: Colors.grey.shade600,
                  ),
                ),
                const SizedBox(height: 2),
                Text(
                  item.value,
                  style: const TextStyle(
                    fontSize: 14,
                    fontWeight: FontWeight.w500,
                    color: Colors.black87,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  _StatusInfo _getStatusInfo() {
    switch (transcription.status.toLowerCase()) {
      case 'completed':
        return _StatusInfo(
          icon: Icons.check_circle,
          color: Colors.green,
          title: 'Completed Successfully',
          subtitle: 'Your audio has been transcribed',
        );
      case 'pending':
        return _StatusInfo(
          icon: Icons.pending,
          color: Colors.orange,
          title: 'Processing',
          subtitle: 'Transcription in progress',
        );
      case 'failed':
        return _StatusInfo(
          icon: Icons.error,
          color: Colors.red,
          title: 'Failed',
          subtitle: 'Transcription could not be completed',
        );
      default:
        return _StatusInfo(
          icon: Icons.info,
          color: Colors.blue,
          title: transcription.status,
          subtitle: 'Status information',
        );
    }
  }

  String _formatDate(DateTime date) {
    return DateFormat('MMM dd, yyyy • HH:mm').format(date);
  }

  String _formatDuration(double duration) {
    if (duration <= 0) return 'N/A';
    final seconds = duration.toInt();
    final minutes = seconds ~/ 60;
    final remainingSeconds = seconds % 60;
    if (minutes > 0) {
      return '${minutes}m ${remainingSeconds}s';
    }
    return '${seconds}s';
  }

  String _formatId(String id) {
    if (id.length > 8) {
      return '${id.substring(0, 8)}...';
    }
    return id;
  }

  void _copyToClipboard(BuildContext context) {
    Clipboard.setData(ClipboardData(text: transcription.text));
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(
        content: Text('Copied to clipboard'),
        duration: Duration(seconds: 2),
        behavior: SnackBarBehavior.floating,
      ),
    );
  }
}

class _StatusInfo {
  final IconData icon;
  final Color color;
  final String title;
  final String subtitle;

  _StatusInfo({
    required this.icon,
    required this.color,
    required this.title,
    required this.subtitle,
  });
}

class _MetadataItem {
  final IconData icon;
  final String label;
  final String value;

  _MetadataItem({
    required this.icon,
    required this.label,
    required this.value,
  });
}

