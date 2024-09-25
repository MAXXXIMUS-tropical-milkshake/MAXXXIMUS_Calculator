import 'package:bloc/bloc.dart';
import 'package:ui/calculator/calculator_client.dart';

class HistoryCubit extends Cubit<List<String>> {
  HistoryCubit(List<String> init) : super(init);

  Future<void> fetchHistory() async {
    emit(await CalculatorClient.fetchHistory());
  }
}
