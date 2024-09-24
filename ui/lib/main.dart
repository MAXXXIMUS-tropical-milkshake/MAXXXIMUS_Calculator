import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/app.dart';
import 'package:ui/calculator_observer.dart';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:window_size/window_size.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  Bloc.observer = const CalculatorObserver();
  if (!kIsWeb && (Platform.isWindows || Platform.isLinux || Platform.isMacOS)) {
    setWindowTitle('Maxxximus Calculator');

    setWindowMinSize(const Size(600, 800));
    setWindowMaxSize(const Size(600, 800));

    setWindowFrame(const Rect.fromLTWH(100, 100, 600, 800));
  }
  runApp(CalculatorApp());
}
