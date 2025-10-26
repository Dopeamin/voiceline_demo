import 'package:flutter/foundation.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:voiceline/core/api/api_client.dart';

class AuthContext extends ChangeNotifier {
  final FlutterSecureStorage _storage = const FlutterSecureStorage();
  final ApiClient _apiClient = ApiClient();
  
  String? _token;
  Map<String, dynamic>? _user;
  bool _isLoading = true;

  String? get token => _token;
  Map<String, dynamic>? get user => _user;
  bool get isAuthenticated => _token != null;
  bool get isLoading => _isLoading;

  AuthContext() {
    // Set up auto-logout on 401 errors
    _apiClient.setOnUnauthorizedCallback(() {
      logout();
    });
    
    _loadAndValidateToken();
  }

  Future<void> _loadAndValidateToken() async {
    try {
      _token = await _storage.read(key: 'auth_token');
      final userJson = await _storage.read(key: 'user_data');
      
      if (_token != null) {
        // Validate token with backend
        final isValid = await _validateToken(_token!);
        
        if (!isValid) {
          // Token is invalid or expired, clear it
          debugPrint('Token validation failed - logging out');
          await _clearStorage();
          _token = null;
          _user = null;
        } else if (userJson != null) {
          // Token is valid, keep user data
          debugPrint('Token validated successfully');
        }
      }
    } catch (e) {
      debugPrint('Error loading token: $e');
      // On error, clear everything to be safe
      await _clearStorage();
      _token = null;
      _user = null;
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  Future<bool> _validateToken(String token) async {
    try {
      _apiClient.setAuthToken(token);
      
      // Try to get transcriptions as a way to validate the token
      // This endpoint requires authentication, so it will fail if token is invalid
      final response = await _apiClient.get('/transcriptions');
      
      return response.statusCode == 200;
    } catch (e) {
      debugPrint('Token validation error: $e');
      return false;
    }
  }

  Future<void> _clearStorage() async {
    await _storage.delete(key: 'auth_token');
    await _storage.delete(key: 'user_data');
  }

  Future<void> login(String token, Map<String, dynamic> user) async {
    _token = token;
    _user = user;
    
    await _storage.write(key: 'auth_token', value: token);
    
    notifyListeners();
  }

  Future<void> logout() async {
    _token = null;
    _user = null;
    
    await _clearStorage();
    _apiClient.clearAuthToken();
    
    notifyListeners();
  }
}

