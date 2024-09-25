import 'dart:convert';

import 'package:http/http.dart' as http;

class Auth {
  static String _token = "";

  static Future<String> get token async {
    if (_token != "") return _token;
    _token = await CalculatorClient.auth();
    return _token;
  }
}

class CalculatorClient {
  static Future<String> auth() async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/auth/signup');
    final response = await client.get(url);
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return decodedResponse["status"] == "success"
        ? decodedResponse["data"]
        : "";
  }

  static Future<int?> calculate(String expr) async {
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/calculate', {"data": expr});
    final response = await client.get(url);
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return decodedResponse["status"] == "success"
        ? decodedResponse["data"]
        : null;
  }

  static Future<List<String>> fetchHistory() async {
    return ["1", "2 + 3", "5 * 6 * 7"];
    final client = http.Client();
    final url = Uri.http('localhost:8080', 'api/v1/history', {"limit": "10"});
    // ignore: avoid_print
    print("sent history request");
    final response =
        await client.get(url, headers: {"Authorization": await Auth.token});
    // ignore: avoid_print
    print("got history response");
    final decodedResponse = jsonDecode(utf8.decode(response.bodyBytes)) as Map;
    client.close();
    return decodedResponse["status"] == "success"
        ? (decodedResponse["data"]["history"] as List<Map>)
            .map((e) => e["expression"] as String)
            .toList()
        : [];
  }
}
