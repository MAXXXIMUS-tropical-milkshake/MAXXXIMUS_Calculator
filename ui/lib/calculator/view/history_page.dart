import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/calculator/cubit/history_cubit.dart';
import 'package:ui/calculator/view/calculator_buttons.dart';

class HistoryPage extends StatelessWidget {
  const HistoryPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => HistoryCubit(),
      child: const HistoryView(),
    );
  }
}

class HistoryView extends StatelessWidget {
  const HistoryView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: const Color.fromARGB(196, 81, 77, 77),
        appBar: AppBar(
          title: const CalculatorCalcButton(),
          backgroundColor: Colors.transparent,
        ),
        body: Center(
            child: BlocBuilder<HistoryCubit, List<String>>(
                builder: (ctx, state) => const Column(
                      children: [],
                    ))));
  }
}
