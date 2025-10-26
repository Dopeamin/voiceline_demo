import 'package:dio/dio.dart';
import 'app_exception.dart';

class ErrorHandler {
  static AppException handleError(dynamic error) {
    if (error is DioException) {
      return _handleDioError(error);
    } else if (error is AppException) {
      return error;
    } else {
      return AppException(
        message: error.toString(),
        code: 'UNKNOWN_ERROR',
      );
    }
  }

  static AppException _handleDioError(DioException error) {
    switch (error.type) {
      case DioExceptionType.connectionTimeout:
      case DioExceptionType.sendTimeout:
      case DioExceptionType.receiveTimeout:
        return TimeoutException();

      case DioExceptionType.connectionError:
        return NetworkException(
          message: 'Unable to connect to server. Please check your connection.',
        );

      case DioExceptionType.badResponse:
        return _handleResponseError(error.response);

      case DioExceptionType.cancel:
        return AppException(
          message: 'Request cancelled',
          code: 'CANCELLED',
        );

      default:
        return NetworkException();
    }
  }

  static AppException _handleResponseError(Response? response) {
    if (response == null) {
      return ServerException();
    }

    final statusCode = response.statusCode;
    final data = response.data;

    // Extract message from response
    String? message;
    String? code;

    if (data is Map<String, dynamic>) {
      message = data['message'] as String?;
      code = data['code'] as String?;
    }

    switch (statusCode) {
      case 400:
        return ValidationException(message: message ?? 'Invalid request');

      case 401:
        return UnauthorizedException(message: message ?? 'Please login again');

      case 403:
        return AppException(
          message: message ?? 'Access denied',
          code: code ?? 'FORBIDDEN',
          statusCode: 403,
        );

      case 404:
        return AppException(
          message: message ?? 'Resource not found',
          code: code ?? 'NOT_FOUND',
          statusCode: 404,
        );

      case 409:
        return ConflictException(message: message);

      case 422:
        return ValidationException(message: message ?? 'Validation failed');

      case 500:
      case 502:
      case 503:
      case 504:
        return ServerException(message: message);

      default:
        return AppException(
          message: message ?? 'Something went wrong',
          code: code ?? 'UNKNOWN',
          statusCode: statusCode,
        );
    }
  }

  static String getErrorMessage(dynamic error) {
    final appException = handleError(error);
    return appException.message;
  }
}

