import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:voiceline/core/context/auth_context.dart';
import 'package:voiceline/core/navigation/app_navigator.dart';

void main() {
  runApp(const VoicelineApp());
}

class VoicelineApp extends StatelessWidget {
  const VoicelineApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => AuthContext()),
      ],
      child: MaterialApp(
        title: 'Voiceline',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
          useMaterial3: true,
        ),
        home: const AppNavigator(),
      ),
    );
  }
}

