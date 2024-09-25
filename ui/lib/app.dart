import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/calculator/view/calculator_page.dart';
import 'package:ui/calculator/view/history_page.dart';

class CalculatorApp extends MaterialApp {
  CalculatorApp({super.key}) : super.router(routerConfig: _router);
}

final GoRouter _router = GoRouter(
  routes: <RouteBase>[
    GoRoute(
      path: '/',
      builder: (BuildContext context, GoRouterState state) {
        // TODO: call history fetch here mb?
        return const CalculatorPage();
      },
      routes: <RouteBase>[
        GoRoute(
          path: 'history',
          builder: (BuildContext context, GoRouterState state) {
            return const HistoryPage();
          },
        ),
      ],
    ),
  ],
);
