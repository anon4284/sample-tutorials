import 'package:angular2/core.dart';
import 'dart:html';
import 'dart:async';

@Injectable()
class FlashMsgService {
  static void msg(Map result) {
    if(result["Valid"]) {
      querySelector('.alert-success').text = result["Msg"];
      querySelector('.alert-success').style.display = "block";
      new Timer(new Duration(seconds: 1), () =>
         querySelector('.alert-success').style.display = "none");

    } else {
      querySelector('.alert-danger').text = result["Msg"];
      querySelector('.alert-danger').style.display = "block";
      new Timer(new Duration(seconds: 1), () =>
         querySelector('.alert-danger').style.display = "none");
    }
  }
}
