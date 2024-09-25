import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/calculator/cubit/calculator_cubit.dart';

class CalculatorButtonStyle {
  static final _borderShape = RoundedRectangleBorder(
    borderRadius: BorderRadius.circular(50),
  );

  static const _padding = EdgeInsets.symmetric(vertical: 24, horizontal: 16);
  static const _textStyle = TextStyle(fontSize: 24);

  static ButtonStyle get digit => ElevatedButton.styleFrom(
        padding: _padding,
        textStyle: _textStyle,
        backgroundColor: const Color.fromARGB(255, 97, 97, 97),
        foregroundColor: Colors.white,
        shape: _borderShape,
      );

  static ButtonStyle get equals => ElevatedButton.styleFrom(
        padding: _padding,
        textStyle: _textStyle,
        backgroundColor: Colors.orange,
        foregroundColor: Colors.white,
        shape: _borderShape,
      );

  static ButtonStyle get clear => ElevatedButton.styleFrom(
        padding: _padding,
        textStyle: _textStyle,
        backgroundColor: Colors.red,
        foregroundColor: Colors.white,
        shape: _borderShape,
      );

  static ButtonStyle get operation => ElevatedButton.styleFrom(
        padding: _padding,
        textStyle: _textStyle,
        backgroundColor: const Color.fromARGB(255, 26, 23, 23),
        foregroundColor: Colors.white,
        shape: _borderShape,
      );

  static ButtonStyle get history => ElevatedButton.styleFrom(
        padding: _padding,
        textStyle: _textStyle,
        backgroundColor: const Color.fromARGB(255, 26, 23, 23),
        foregroundColor: Colors.white,
        shape: _borderShape,
      );
}

class CalculatorClearButton extends StatelessWidget {
  const CalculatorClearButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: CalculatorButtonStyle.clear,
          onPressed: () => context.read<CalculatorCubit>().eraseAll(),
          child: const Text("С"),
        ),
      ),
    );
  }
}

class CalculatorEqualsButton extends StatelessWidget {
  const CalculatorEqualsButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: CalculatorButtonStyle.equals,
          onPressed: () => context.read<CalculatorCubit>().evaluate(),
          child: const Text("="),
        ),
      ),
    );
  }
}

class CalculatorOperationButton extends StatelessWidget {
  const CalculatorOperationButton({
    super.key,
    required this.operation,
  });

  final String operation;

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: CalculatorButtonStyle.operation,
          onPressed: () => context.read<CalculatorCubit>().insertOp(operation),
          child: Text(operation),
        ),
      ),
    );
  }
}

class CalculatorSymButton extends StatelessWidget {
  const CalculatorSymButton({
    super.key,
    required this.sym,
  });

  final String sym;

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: CalculatorButtonStyle.operation,
          onPressed: () => context.read<CalculatorCubit>().insertSym(sym),
          child: Text(sym),
        ),
      ),
    );
  }
}

class CalculatorHistoryButton extends StatelessWidget {
  const CalculatorHistoryButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Align(
        alignment: Alignment.center,
        child: ElevatedButton(
          style: CalculatorButtonStyle.history,
          onPressed: () => context.go("/history"),
          child: const Text("History"),
        ));
  }
}

class CalculatorDigitButton extends StatelessWidget {
  const CalculatorDigitButton({super.key, required this.digit});

  final String digit;

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: CalculatorButtonStyle.digit,
          onPressed: () => context.read<CalculatorCubit>().insertDigit(digit),
          child: Text(digit),
        ),
      ),
    );
  }
}
