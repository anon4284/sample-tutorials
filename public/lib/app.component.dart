import "package:angular2/core.dart";
import 'package:angular2/router.dart';
import 'user/components/login.component.dart';
import 'user/components/signup.component.dart';
import "home/home.component.dart";
import "components/navbar.component.dart";
import "portfolio/portfolio.component.dart";
import "portfolio/components/addProject.component.dart";
import "user/components/userList.component.dart";
import "user/components/profile.component.dart";
import "user/services/user.service.dart";


@Component(
  selector: 'st-app',
  templateUrl:'app.component.html',
  styleUrls: const ["app.component.css"],
  directives: const [ROUTER_DIRECTIVES, NavbarComponent]
)
@RouteConfig(const [const Route(path: '/', component: HomeComponent, name: "Home"),
                    const Route(path: '/login', component: LoginComponent, name: "Login"),
                    const Route(path: '/signup', component: SignupComponent, name: "Signup"),
                    const Route(path: '/addProject', component: AddProjectComponent, name: 'AddProject'),
                    const Route(path: '/projects', component: PortfolioComponent, name: "Portfolio"),
                    const Route(path: '/users', component: UserListComponent, name: "Users"),
                  const Route(path: '/profile', component: ProfileComponent, name: "Profile")])
class AppComponent {
  Router router;

  bool isLoggedIn;

  AppComponent(this.router) {
    router.subscribe((onNext) {
      isLoggedIn = UserService.isLoggedIn();
    });
  }

}
