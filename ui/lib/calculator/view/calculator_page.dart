import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
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
  Widget build(BuildContext context) {
    final textTheme = Theme.of(context).textTheme;
    return Scaffold(
      body: Center(
        child: BlocBuilder<CalculatorCubit, String>(
          builder: (context, state) => Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Text(state, style: textTheme.displayMedium),
              Row(
                children: [
                  for (var i in ["-", "+", "*"])
                    TextButton(
                        onPressed: () =>
                            context.read<CalculatorCubit>().insertOp(i),
                        child: Text(i)),
                ],
              ),
              Row(
                children: [
                  for (var i in [7, 8, 9])
                    TextButton(
                        onPressed: () => context
                            .read<CalculatorCubit>()
                            .insertSymbol(i.toString()),
                        child: Text(i.toString())),
                ],
              ),
              Row(
                children: [
                  for (var i in [4, 5, 6])
                    TextButton(
                        onPressed: () => context
                            .read<CalculatorCubit>()
                            .insertSymbol(i.toString()),
                        child: Text(i.toString()))
                ],
              ),
              Row(
                children: [
                  for (var i in [1, 2, 3])
                    TextButton(
                        onPressed: () => context
                            .read<CalculatorCubit>()
                            .insertSymbol(i.toString()),
                        child: Text(i.toString()))
                ],
              ),
              TextButton(
                onPressed: context.read<CalculatorCubit>().evaluate,
                child: const Text("="),
              ),
              TextButton(
                  onPressed: context.read<CalculatorCubit>().eraseLast,
                  child: const Text("<"))
            ],
          ),
        ),
      ),
    );
  }
}
