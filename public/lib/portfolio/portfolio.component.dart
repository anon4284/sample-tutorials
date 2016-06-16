import "package:angular2/core.dart";
import 'package:angular2/router.dart';
import 'components/empty.component.dart';
import 'components/project.component.dart';


@Component(
  selector: 'st-portfolio',
  templateUrl: 'portfolio.component.html',
  directives: const [ROUTER_DIRECTIVES]
)
@RouteConfig(const [    const Route(path: '/:id', component: ProjectComponent, name: 'Project'),
                        const Route(path: '/', component: EmptyComponent, useAsDefault: true, name:'Empty'),
                      ])
class PortfolioComponent {
  var values = [1,2,3,4,5,6];

  bool isPortfolio() {
    return Uri.base.path == '/projects';
  }


}
