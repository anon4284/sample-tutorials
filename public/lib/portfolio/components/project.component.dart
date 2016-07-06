import "package:angular2/core.dart";
import "../helper/project.dart";
import "package:sampleTutorials/user/services/user.service.dart";
import "dart:html";
import "package:sampleTutorials/comments/comments.component.dart";

@Component(
  selector: 'portfolio-project',
  templateUrl: 'project.component.html',
  styleUrls: const ['project.component.css'],
  directives: const [CommentsComponent]
)
class ProjectComponent implements OnInit {
  Project project;


  @override
  ngOnInit() {
    var resp = UserService.get('/api/projects/getByID?'+Uri.base.query);
    print(resp);
    project = new Project(resp["Item"]["ProjectID"],
    resp["Item"]["Title"],
    resp["Item"]["Description"],
    resp["Item"]["Content"]);

    final NodeValidatorBuilder _htmlValidator=new NodeValidatorBuilder.common()
    ..allowElement('IMG', attributes: ['src','style']);


    querySelector("#content").setInnerHtml(project.content, validator: _htmlValidator);

  }
}
