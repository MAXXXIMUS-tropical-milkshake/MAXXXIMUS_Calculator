import 'package:bloc/bloc.dart';
import 'package:ui/calculator/calculator_client.dart';

const _supportedOps = ["-", "+", "×", "÷"];
const _supportedNonOps = [
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

bool lastisop = false;

class CalculatorCubit extends Cubit<String> {
  CalculatorCubit() : super("");

  void insertSymbol(String symbol) {
    if (!_supportedNonOps.contains(symbol)) return;
    emit(state + symbol);
    lastisop = false;
  }

  void insertOp(String op) {
    if (!_supportedOps.contains(op)) return;
    if (state.isEmpty) return;
    if (lastisop) {
      emit(state.substring(0, state.length - 1) + op);
      return;
    }
    emit(state + op);
    lastisop = true;
  }

  void eraseLast() {
    emit(state.isEmpty ? state : state.substring(0, state.length - 1));
    lastisop = false;
  }

  void eraseAll() {
    emit("");
    lastisop = false;
  }

  Future<void> evaluate() async {
    if (state.isEmpty) return;
    if (lastisop) return;
    emit(await CalculatorClient.calculate(state));
    lastisop = false;
  }
}
