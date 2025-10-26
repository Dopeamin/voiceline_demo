import 'package:flutter/material.dart';
import 'package:voiceline/features/auth/screens/login_screen.dart';
import 'package:voiceline/features/auth/screens/register_screen.dart';

class AuthScreen extends StatefulWidget {
  const AuthScreen({super.key});

  @override
  State<AuthScreen> createState() => _AuthScreenState();
}

class _AuthScreenState extends State<AuthScreen> {
  bool _isLogin = true;

  void _toggleMode() {
    setState(() {
      _isLogin = !_isLogin;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: AnimatedSwitcher(
        duration: const Duration(milliseconds: 300),
        child: _isLogin
            ? LoginScreen(key: const ValueKey('login'), onToggle: _toggleMode)
            : RegisterScreen(key: const ValueKey('register'), onToggle: _toggleMode),
      ),
    );
  }
}

