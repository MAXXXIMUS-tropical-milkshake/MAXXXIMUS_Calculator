import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/calculator/cubit/history_cubit.dart';

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
  Widget build(BuildContext ctx) {
    final textTheme = Theme.of(ctx).textTheme;

    return Scaffold(
        body: Center(
            child: BlocBuilder<HistoryCubit, List<String>>(
                builder: (context, state) => Column(
                      children: [
                        TextButton(
                            onPressed: () async {
                              await context.read<HistoryCubit>().fetchHistory();
                            },
                            child: Text("load history")),
                        TextButton(
                            onPressed: () async => context.go("/"),
                            child: Text("to calculator")),
                        for (var line in state) Text(line)
                      ],
                    ))));
  }
}
