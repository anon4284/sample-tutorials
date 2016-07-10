import "package:angular2/core.dart";
import "dart:html";
import "../services/user.service.dart";
import "../services/flashMsg.service.dart";
import "package:angular2/router.dart";

@Component(
  selector: 'user-login',
  templateUrl: 'login.component.html'
)
class LoginComponent {
  final Router router;

  String useremail;
  String password;

  LoginComponent(this.router){}

  void login() {
    var data = new Map<String,String>();
    data["Useremail"] = useremail;
    data["Password"] = password;

    var resp = UserService.post("/api/login",data);
    FlashMsgService.msg(resp);

    if (resp["Valid"]) {
      window.localStorage["token"] = resp["Token"];
      window.localStorage["userID"] = resp["UserID"];
      window.localStorage["isAdmin"] = resp["IsAdmin"];
      this.router.parent.navigate(["./Login"]);
    }
  }




}
