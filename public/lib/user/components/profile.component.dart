import "package:angular2/core.dart";
import "dart:html";
import "../services/user.service.dart";
import "../services/flashMsg.service.dart";

@Component(
  selector: 'user-profile',
  templateUrl: 'profile.component.html'
)
class ProfileComponent implements OnInit {
  String passwordOld;
  String passwordNew;
  String passwordNewRepeat;

  String email;
  String emailRepeat;

  String username;
  String useremail;
  String profilePicture;

  void getProfileInfo() {
    var resp = UserService.getLoggedIn("/api/user/me");
    print(resp);
    username = resp["Username"];
    useremail = resp["Useremail"];
    profilePicture = resp["ProfilePicture"];
  }


  void changeEmail() {
    var data = new Map<String, String>();
    data["Email"] = email;
    data["EmailRepeat"] = emailRepeat;
  }


  void changePassword() {
    var data = new Map<String, String>();
    data["Old"] = passwordOld;
    data["New"] = passwordNew;
    data["NewRepeat"] = passwordNewRepeat;
    var resp = UserService.postLoggedIn("/api/user/password/change", data);
    FlashMsgService.msg(resp);
  }

  @override
  ngOnInit() {
      getProfileInfo();
  }
}
