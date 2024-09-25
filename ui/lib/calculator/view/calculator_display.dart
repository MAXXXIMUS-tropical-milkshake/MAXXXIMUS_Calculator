import 'package:flutter/material.dart';

class CalculatorDisplay extends StatelessWidget {
  const CalculatorDisplay({
    super.key,
    required this.displayState,
    required this.textTheme,
  });

  final String displayState;
  final TextTheme textTheme;

  @override
  Widget build(BuildContext context) {
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
            displayState,
            style: textTheme.displayMedium
                ?.copyWith(fontSize: 48, color: Colors.white),
            textAlign: TextAlign.right,
          ),
        ),
      ),
    );
  }
}
