import "package:angular2/core.dart";
import "package:sampleTutorials/user/services/user.service.dart";
import 'package:sampleTutorials/user/services/flashMsg.service.dart';

@Component(
  selector: 'user-signup',
  templateUrl: 'userList.component.html',
  styleUrls: const ["userList.component.css"]
)
class UserListComponent implements OnInit {
  List<UserInfo> users;
  bool isAdmin;

  void delete(String userID) {
    var map = new Map<String,String>();
    map["UserID"] = userID;
    var resp = UserService.postLoggedIn("/api/admin/user/delete", map);
    FlashMsgService.msg(resp);
    getUsers();
  }


  void getUsers() {
    var limit = "10";
    var url = "/api/user/getAll?limit="+limit;
    var resp = UserService.getLoggedIn(url);
    print(resp);
    users = new List<UserInfo>();
    for(int i = 0; i < resp["Users"].length; i++ ) {
        var user = resp["Users"][i];
        users.add(new UserInfo(user["UserID"], user["Username"], user["ProfilePicture"]));
    }
  }


  @override
  ngOnInit() {
    getUsers();
    isAdmin = UserService.isAdmin();
  }
}


class UserInfo {
  final String userID;
  final String username;
  final String profilePicture;
  UserInfo(this.userID, this.username, this.profilePicture);
}
