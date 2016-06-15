import "package:angular2/core.dart";
import "dart:html";
import 'dart:convert';


@Injectable()
class UserService {
  static Map post(String url, Map data) {
    HttpRequest request = new HttpRequest();
    request.open("POST", url, async: false);
    var jsonData = JSON.encode(data);
    request.send(jsonData);
    Map resp = JSON.decode(request.response.toString());
    return resp;
  }

  static void getLoggedIn(String url, Map data) {
    HttpRequest.request(url, method: "GET", requestHeaders: {
      'x-access-userid': window.localStorage["userID"],
      'x-access-token': window.localStorage["token"],
     });
  }

}
