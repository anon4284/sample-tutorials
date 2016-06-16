import "package:angular2/core.dart";
import 'package:angular2/router.dart';
import 'user/components/login.component.dart';
import 'user/components/signup.component.dart';
import "home/home.component.dart";
import "components/navbar.component.dart";
import "portfolio/portfolio.component.dart";
import "portfolio/components/addProject.component.dart";


@Component(
  selector: 'st-app',
  templateUrl:'app.component.html',
  directives: const [ROUTER_DIRECTIVES, NavbarComponent]
)
@RouteConfig(const [const Route(path: '/', component: HomeComponent, name: "Home"),
                    const Route(path: '/login', component: LoginComponent, name: "Login"),
                    const Route(path: '/signup', component: SignupComponent),
                    const Route(path: '/addProject', component: AddProjectComponent, name: 'AddProject'),
                    const Route(path: '/projects/...', component: PortfolioComponent, name: "Portfolio")])
class AppComponent implements OnInit {

  @override
  ngOnInit() {
    // TODO: implement ngOnInit
  }
}
