import 'dart:convert';

import 'package:http/http.dart' as http;

class CalculatorClient {
  Future<String> calculate(String expr) async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'handle');
    final response = await client.post(url, body: expr);
    final decodedResponse =
        jsonDecode(utf8.decode(response.bodyBytes)) as String;
    client.close();
    return decodedResponse;
  }
}
