import 'dart:convert';

import 'package:http/http.dart' as http;

class CalculatorClient {
  // TODO: may get non-string as argument
  Future<String> calculate(String expr) async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/calculate', {"data": expr});
    final response = await client.get(url);
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return decodedResponse["data"]?.toString() ?? ":(";
  }
}
