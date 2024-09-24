import 'package:bloc/bloc.dart';
import 'package:ui/calculator/calculator_client.dart';

const _supportedOps = [
  "+",
  "-",
  "*",
  "/",
];
const _supportedNonOps = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "00", "."];

class CalculatorCubit extends Cubit<String> {
  CalculatorCubit() : super("");

  void insertSymbol(String symbol) {
    if (!_supportedNonOps.contains(symbol)) return;
    emit(state + symbol);
  }

  void insertOp(String op) {
    if (!_supportedOps.contains(op)) return;
    emit(state + op);
  }

  void eraseLast() =>
      emit(state.isEmpty ? state : state.substring(0, state.length - 1));

  void eraseAll() =>
      emit("");

  Future<void> evaluate() async {
    emit(await CalculatorClient.calculate(state));
  }
}
