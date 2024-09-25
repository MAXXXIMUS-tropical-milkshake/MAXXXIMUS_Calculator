import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/calculator/view/calculator_page.dart';
import 'package:ui/calculator/view/history_page.dart';

class CalculatorApp extends MaterialApp {
  CalculatorApp({super.key}) : super.router(routerConfig: _router);
}

final GoRouter _router = GoRouter(
  initialLocation: "/calculator",
  routes: <RouteBase>[
    GoRoute(
      path: '/calculator',
      builder: (BuildContext context, GoRouterState state) {
        return const CalculatorPage();
      },
    ),
    GoRoute(
      path: '/history',
      builder: (BuildContext context, GoRouterState state) {
        return const HistoryPage();
      },
    ),
  ],
);
