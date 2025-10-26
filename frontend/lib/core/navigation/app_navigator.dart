import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:voiceline/core/context/auth_context.dart';
import 'package:voiceline/features/auth/screens/auth_screen.dart';
import 'package:voiceline/features/recording/screens/recording_screen.dart';

class AppNavigator extends StatelessWidget {
  const AppNavigator({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<AuthContext>(
      builder: (context, authContext, _) {
        if (authContext.isLoading) {
          return Scaffold(
            body: Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const CircularProgressIndicator(),
                  const SizedBox(height: 16),
                  Text(
                    'Validating session...',
                    style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                          color: Colors.grey,
                        ),
                  ),
                ],
              ),
            ),
          );
        }

        if (authContext.isAuthenticated) {
          return const RecordingScreen();
        }

        return const AuthScreen();
      },
    );
  }
}

