import "package:angular2/core.dart";
import "../services/user.service.dart";
import '../services/flashMsg.service.dart';

@Component(
  selector: 'user-signup',
  templateUrl: 'signup.component.html'
)
class SignupComponent {
  String useremail;
  String username;
  String password;
  String passwordRepeat;

  void signup() {
    if(password == passwordRepeat) {
      var data = new Map<String,String>();
      data["Useremail"] = useremail;
      data["Username"] = useremail;
      data["Password"] = password;
      var resp = UserService.post("/api/user/signup",data);
      FlashMsgService.msg(resp);
      print(resp);
    } else {
      var map = new Map<String, String>();
      map["Valid"] = false;
      map["Msg"] = "Passwords dont match";
      FlashMsgService.msg(map);
    }

  }

}
