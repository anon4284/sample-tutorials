import "package:angular2/core.dart";
import "dart:html";
import "../services/user.service.dart";
import "../services/flashMsg.service.dart";

@Component(
  selector: 'user-profile',
  templateUrl: 'profile.component.html'
)
class ProfileComponent {
  String useremail;
  String password;

  void changePassword() {
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
