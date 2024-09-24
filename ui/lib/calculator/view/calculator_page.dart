import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/calculator/cubit/calculator_cubit.dart';

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
  Widget build(BuildContext ctx) {
    final textTheme = Theme.of(ctx).textTheme;
    final screenSize = MediaQuery.of(ctx).size;

    final buttonStyle = ElevatedButton.styleFrom(
      padding: const EdgeInsets.symmetric(vertical: 24.0),
      textStyle: const TextStyle(fontSize: 24),
      backgroundColor: const Color.fromARGB(255, 97, 97, 97),
      foregroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(50),
      ),
    );

    final buttonStyleOrange = ElevatedButton.styleFrom(
      padding: const EdgeInsets.symmetric(vertical: 24.0),
      textStyle: const TextStyle(fontSize: 24),
      backgroundColor: Colors.orange,
      foregroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(50),
      ),
    );

    final buttonStyleRed = ElevatedButton.styleFrom(
      padding: const EdgeInsets.symmetric(vertical: 24.0),
      textStyle: const TextStyle(fontSize: 24),
      backgroundColor: Colors.red,
      foregroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(50),
      ),
    );

    final buttonStyleGrey = ElevatedButton.styleFrom(
      padding: const EdgeInsets.symmetric(vertical: 24.0),
      textStyle: const TextStyle(fontSize: 24),
      backgroundColor: const Color.fromARGB(255, 26, 23, 23),
      foregroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(50),
      ),
    );

    final buttonStyleHistory = ElevatedButton.styleFrom(
      padding: const EdgeInsets.all(16),
      textStyle: const TextStyle(fontSize: 20),
      backgroundColor: const Color.fromARGB(255, 26, 23, 23),
      foregroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(50),
      ),
    );

    return Scaffold(
      backgroundColor: const Color.fromARGB(196, 81, 77, 77),
      body: BlocBuilder<CalculatorCubit, String>(
        builder: (context, state) {
          final displayState = state.replaceAll('*', '×').replaceAll('/', '÷');

          return Container(
            width: screenSize.width,
            height: screenSize.height,
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Align(
                    alignment: Alignment.center,
                    child: ElevatedButton(
                      style: buttonStyleHistory,
                      onPressed: () => context.go("/history"),
                      child: const Text("History"),
                    )),
                const SizedBox(height: 10),
                Expanded(
                  flex: 2,
                  child: Container(
                    padding: const EdgeInsets.all(16.0),
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(12),
                      color: const Color.fromARGB(255, 97, 97, 97),
                    ),
                    child: Align(
                      alignment: Alignment.centerRight,
                      child: Text(
                        displayState,
                        style: textTheme.displayMedium
                            ?.copyWith(fontSize: 48, color: Colors.white),
                        textAlign: TextAlign.right,
                      ),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var i in ["-", "+", "×", "÷"])
                        Expanded(
                          child: Padding(
                            padding: const EdgeInsets.all(8.0),
                            child: ElevatedButton(
                              style: buttonStyleGrey,
                              onPressed: () => context
                                  .read<CalculatorCubit>()
                                  .insertOp(
                                      i == "×" ? "*" : (i == "÷" ? "/" : i)),
                              child: Text(i),
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
                      for (var i in [7, 8, 9])
                        CalculatorButton(
                            buttonStyle: buttonStyle, symbol: i.toString()),
                      Expanded(
                        child: Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                            style: buttonStyleOrange,
                            onPressed: () =>
                                context.read<CalculatorCubit>().evaluate(),
                            child: const Text("="),
                          ),
                        ),
                      )
                    ],
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      for (var i in [4, 5, 6])
                        CalculatorButton(
                            buttonStyle: buttonStyle, symbol: i.toString()),
                      Expanded(
                        child: Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                            style: buttonStyleGrey,
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
                      for (var i in [1, 2, 3])
                        CalculatorButton(
                            buttonStyle: buttonStyle, symbol: i.toString()),
                      Expanded(
                        child: Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                            style: buttonStyleRed,
                            onPressed: () =>
                                context.read<CalculatorCubit>().eraseAll(),
                            child: const Text("С"),
                          ),
                        ),
                      )
                    ],
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      CalculatorButton(buttonStyle: buttonStyle, symbol: '0'),
                      CalculatorButton(buttonStyle: buttonStyleGrey, symbol: '.'),
                      CalculatorButton(buttonStyle: buttonStyleGrey, symbol: '('),
                      CalculatorButton(buttonStyle: buttonStyleGrey, symbol: ')'),
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

class CalculatorButton extends StatelessWidget {
  const CalculatorButton(
      {super.key, required this.buttonStyle, required this.symbol});

  final ButtonStyle buttonStyle;
  final String symbol;

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ElevatedButton(
          style: buttonStyle,
          onPressed: () => context.read<CalculatorCubit>().insertSymbol(symbol),
          child: Text(symbol),
        ),
      ),
    );
  }
}
