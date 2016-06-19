import "package:angular2/core.dart";
import 'components/project.component.dart';
import 'package:angular2/router.dart';
import 'package:sampleTutorials/user/services/user.service.dart';
import 'helper/project.dart';


@Component(
  selector: 'st-portfolio',
  templateUrl: 'portfolio.component.html',
  directives: const [ProjectComponent]
)
class PortfolioComponent implements OnInit {
  List<Project> projects;
  final Router router;

  PortfolioComponent(this.router) {}

  bool isPortfolio() {
    return Uri.base.path == '/projects' && Uri.base.queryParameters["id"] == null;
  }

  void goToProject(String id) {
    print(id);
    this.router.parent.navigate(['./Portfolio', {"id": id}]);
  }

  @override
  ngOnInit() {
    if(isPortfolio()) {
      var resp = UserService.get('/api/projects/get');
      print(resp);
      projects = new List<Project>();
      for(int i = 0; i < resp["Items"].length; i++ ) {
          print(resp["Items"][i]["Title"]);
        projects.add(new Project(resp["Items"][i]["ProjectID"],resp["Items"][i]["Title"],resp["Items"][i]["Description"],resp["Items"][i]["Content"]));
      }
    }
  }
}
