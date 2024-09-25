import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/calculator/cubit/calculator_cubit.dart';
import 'package:ui/calculator/view/calculator_buttons.dart';
import 'package:ui/calculator/view/calculator_display.dart';

class CalculatorPage extends StatelessWidget {
  const CalculatorPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => CalculatorCubit(),
      child: const CalculatorView(),
    );
  }
}

class CalculatorView extends StatelessWidget {
  const CalculatorView({super.key});

  @override
  // ignore: avoid_renaming_method_parameters
  Widget build(BuildContext ctx) {
    final textTheme = Theme.of(ctx).textTheme;
    final screenSize = MediaQuery.of(ctx).size;

    return Scaffold(
      backgroundColor: const Color.fromARGB(196, 81, 77, 77),
      body: BlocBuilder<CalculatorCubit, String>(
        builder: (context, state) {
          return Container(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                const CalculatorHistoryButton(),
                const SizedBox(height: 10),
                CalculatorDisplay(displayState: state, textTheme: textTheme),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var op in ["-", "+", "×", "÷"])
                        CalculatorOperationButton(operation: op),
                    ],
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var digit in ["7", "8", "9"])
                        CalculatorDigitButton(digit: digit),
                      const CalculatorEqualsButton()
                    ],
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var digit in ["4", "5", "6"])
                        CalculatorDigitButton(digit: digit),
                      Expanded(
                        child: Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                            style: CalculatorButtonStyle.operation,
                            onPressed: () =>
                                context.read<CalculatorCubit>().eraseLast(),
                            child: const Text("←"),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var digit in ["1", "2", "3"])
                        CalculatorDigitButton(digit: digit),
                      const CalculatorClearButton()
                    ],
                  ),
                ),
                const Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      CalculatorDigitButton(digit: '0'),
                      CalculatorSymButton(sym: '.'),
                      CalculatorSymButton(sym: '('),
                      CalculatorSymButton(sym: ')'),
                    ],
                  ),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
