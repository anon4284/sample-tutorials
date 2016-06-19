import "package:angular2/core.dart";
import "dart:html";
import 'dart:convert';


@Injectable()
class UserService {


  static Map get(String url) {
    HttpRequest request = new HttpRequest();
    request.open("GET", url, async: false);
    request.send();
    Map resp = JSON.decode(request.response.toString());
    return resp;
  }


  static Map post(String url, Map data) {
    HttpRequest request = new HttpRequest();
    request.open("POST", url, async: false);
    var jsonData = JSON.encode(data);
    request.send(jsonData);
    Map resp = JSON.decode(request.response.toString());
    return resp;
  }

  static Map getLoggedIn(String url) {
    HttpRequest request = new HttpRequest();
     request.open("GET", url, async: false);
     request.setRequestHeader('x-access-userid', window.localStorage["userID"]);
     request.setRequestHeader('x-access-token', window.localStorage["token"]);
     request.send();
     Map resp = JSON.decode(request.response.toString());
     return resp;
  }

  static Map postLoggedIn(String url, Map data) {
    HttpRequest request = new HttpRequest();
     request.open("POST", url, async: false);
     request.setRequestHeader('x-access-userid', window.localStorage["userID"]);
     request.setRequestHeader('x-access-token', window.localStorage["token"]);
     var jsonData = JSON.encode(data);
     request.send(jsonData);
     Map resp = JSON.decode(request.response.toString());
     return resp;
  }

}
