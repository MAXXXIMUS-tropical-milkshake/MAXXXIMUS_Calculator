import 'dart:convert';

import 'package:http/http.dart' as http;

class CalculatorClient {
  // TODO: may get non-string as argument
  static Future<String> calculate(String expr) async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/calculate', {"data": expr});
    final response = await client.get(url);
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return decodedResponse["data"]?.toString() ?? ":(";
  }

  static Future<List<String>> fetchHistory() async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/history', {"limit": "10"});
    // ignore: avoid_print
    print("sent history request");
    final response = await client.get(url);
    // ignore: avoid_print
    print("got history response");
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return (decodedResponse["data"]["history"] as List<Map>)
        .map((e) => e["expression"] as String)
        .toList();
  }
}
