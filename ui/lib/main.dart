import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/app.dart';
import 'package:ui/calculator_observer.dart';

void main() {
  Bloc.observer = const CalculatorObserver();
  runApp(CalculatorApp());
}
