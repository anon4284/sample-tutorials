import "package:angular2/core.dart";
import "dart:html";
import "package:sampleTutorials/pkg/ckeditor/ckeditor.dart";
import 'package:sampleTutorials/user/services/user.service.dart';
import 'package:sampleTutorials/user/services/flashMsg.service.dart';

@Component(
  selector: 'portfolio-add',
  templateUrl: 'addProject.component.html'
)
class AddProjectComponent implements OnInit {
  String title;
  String description;
  CKEDITOR ck;

  void add() {
    var map = new Map<String, String>();
    map["Title"] = title;
    map["Description"] = description;
    map["Content"] = ck.getData();
    uploadImage();
    var resp = UserService.postLoggedIn('/api/projects/add', map);
    FlashMsgService.msg(resp);
  }

  void uploadImage() {
      FileUploadInputElement input = querySelector("#thumbnail");
      if(input.files.isEmpty)
      return;
      var name = input.files[0].name;
      print(name);
  }

  @override
  ngOnInit() {
    ck = CKEDITOR.replace("content");
  }
}
