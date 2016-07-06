import "package:angular2/core.dart";
import 'package:sampleTutorials/user/services/user.service.dart';
import 'package:sampleTutorials/user/services/flashMsg.service.dart';
import 'package:sampleTutorials/comments/helper/comment.dart';

@Component(
  selector: 'my-comments',
  templateUrl: 'comments.component.html',
  styleUrls: const ['comments.component.css']
)
class CommentsComponent implements OnInit{
  @Input()
  String projectID;

  String content;

  List<Comment> comments;

  void addComment() {
    print(projectID);
    var map = new Map<String, String>();
    map["ProjectID"] = projectID;
    map["Content"] = content;
    var resp = UserService.postLoggedIn('/api/user/comments/add', map);
    FlashMsgService.msg(resp);
    getComments();
  }

  @override
  ngOnInit() {
    getComments();
  }

  void getComments() {
    var resp = UserService.get('/api/comments/getByProjectID?'+Uri.base.query);
    print(resp);
    comments = new List<Comment>();
    for(int i = 0; i < resp["Items"].length; i++ ) {
      comments.add(new Comment(resp["Items"][i]["Username"],resp["Items"][i]["UserID"],resp["Items"][i]["Content"],resp["Items"][i]["CreatedAt"]));
    }
  }
}
