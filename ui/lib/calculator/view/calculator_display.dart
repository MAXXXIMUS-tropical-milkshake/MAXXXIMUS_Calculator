import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/calculator/cubit/calculator_cubit.dart';

class CalculatorDisplay extends StatelessWidget {
  const CalculatorDisplay({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<CalculatorCubit, String>(builder: (context, state) {
      return Expanded(
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
              state,
              style: Theme.of(context)
                  .textTheme
                  .displayMedium
                  ?.copyWith(fontSize: 48, color: Colors.white),
              textAlign: TextAlign.right,
            ),
          ),
        ),
      );
    });
  }
}
