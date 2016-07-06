import "package:angular2/core.dart";
import "dart:html";
import "../services/user.service.dart";
import "../services/flashMsg.service.dart";

@Component(
  selector: 'user-login',
  templateUrl: 'login.component.html'
)
class LoginComponent {
  String useremail = "schnake@gmail.com";
  String password = "Katzenschnaken";

  void login() {
    var data = new Map<String,String>();
    data["Useremail"] = useremail;
    data["Password"] = password;

    var resp = UserService.post("/api/user/login",data);
    FlashMsgService.msg(resp);

    if (resp["Valid"]) {
      window.localStorage["token"] = resp["Token"];
      window.localStorage["userID"] = resp["UserID"];
    }
  }




}
