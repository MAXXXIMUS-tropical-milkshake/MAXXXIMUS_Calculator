import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ui/calculator/calculator_client.dart';
import 'package:ui/calculator/cubit/history_cubit.dart';
import 'package:ui/calculator/view/calculator_buttons.dart';

class HistoryPage extends StatelessWidget {
  const HistoryPage({super.key});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: CalculatorClient.fetchHistory(),
      builder: (context, snapshot) => snapshot.hasData
          ? BlocProvider(
              create: (_) => HistoryCubit(
                  snapshot.hasData ? snapshot.data as List<String> : []),
              child: const HistoryView(),
            )
          : Scaffold(
              body: Container(),
            ),
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
                builder: (ctx, state) => Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        for (final expr in state)
                          Card(
                              color: const Color.fromARGB(255, 26, 23, 23),
                              child: Padding(
                                padding: const EdgeInsets.all(8.0),
                                child: Text(
                                  expr,
                                  style: const TextStyle(
                                      fontSize: 24, color: Colors.white),
                                ),
                              ))
                      ],
                    ))));
  }
}
