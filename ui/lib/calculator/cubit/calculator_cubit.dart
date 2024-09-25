import 'package:bloc/bloc.dart';
import 'package:flutter/material.dart';
import 'package:ui/calculator/calculator_client.dart';

const _supportedOps = ["-", "+", "×", "÷"];
const _supportedSyms = [".", "(", ")"];
const _supportedDigits = [
  "1",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
  "0",
  "0",
  ".",
  "(",
  ")"
];

class CalculatorCubit extends Cubit<String> {
  CalculatorCubit() : super("");

  bool get _lastSymbolIsOp =>
      state.isNotEmpty && !_supportedOps.contains(state[state.length - 1]);

  void insertDigit(String symbol) {
    if (!_supportedDigits.contains(symbol)) return;
    emit(state + symbol);
  }

  void insertOp(String op) {
    if (!_supportedOps.contains(op)) return;
    if (state.isEmpty || !_lastSymbolIsOp) return;
    emit(state + op);
  }

  void insertSym(String sym) {
    if (!_supportedSyms.contains(sym)) return;
    emit(state + sym);
  }

  void eraseLast() {
    emit(state.isEmpty ? state : state.substring(0, state.length - 1));
  }

  void eraseAll() {
    emit("");
  }

  Future<void> evaluate(BuildContext context) async {
    if (state.isEmpty) return;
    final expr = state.replaceAll("×", "*").replaceAll("÷", "/");
    final response = await CalculatorClient.calculate(expr);
    emit(response?.toString() ?? "");
    // ignore: use_build_context_synchronously
    ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text(
      "Error pasrsing expression",
      style: TextStyle(fontSize: 24),
    )));
  }
}
