class AppException implements Exception {
  final String message;
  final String? code;
  final int? statusCode;

  AppException({
    required this.message,
    this.code,
    this.statusCode,
  });

  @override
  String toString() => message;
}

class NetworkException extends AppException {
  NetworkException({String? message})
      : super(
          message: message ?? 'Network error. Please check your connection.',
          code: 'NETWORK_ERROR',
        );
}

class TimeoutException extends AppException {
  TimeoutException()
      : super(
          message: 'Request timed out. Please try again.',
          code: 'TIMEOUT',
        );
}

class UnauthorizedException extends AppException {
  UnauthorizedException({String? message})
      : super(
          message: message ?? 'Unauthorized. Please login again.',
          code: 'UNAUTHORIZED',
          statusCode: 401,
        );
}

class ValidationException extends AppException {
  final Map<String, String>? fieldErrors;

  ValidationException({
    String? message,
    this.fieldErrors,
  }) : super(
          message: message ?? 'Validation failed',
          code: 'VALIDATION_ERROR',
          statusCode: 400,
        );
}

class ServerException extends AppException {
  ServerException({String? message})
      : super(
          message: message ?? 'Server error. Please try again later.',
          code: 'SERVER_ERROR',
        );
}

class ConflictException extends AppException {
  ConflictException({String? message})
      : super(
          message: message ?? 'Resource already exists',
          code: 'CONFLICT',
          statusCode: 409,
        );
}

