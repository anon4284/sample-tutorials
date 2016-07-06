import "package:angular2/core.dart";
import 'package:angular2/router.dart';
import '../user/services/user.service.dart';

@Component(
  selector: 'st-navbar',
  templateUrl: 'navbar.component.html',
  styleUrls: const ['navbar.component.css'],
  directives: const [RouterLink]
)
class NavbarComponent {
  Router router;

  NavbarComponent(Router router) {
    this.router = router;
  }

  bool isLoggedIn = UserService.isLoggedIn();

  void logOut() {
    UserService.logOut();
    isLoggedIn = false;
    this.router.navigate(['/Home']);
  }
}
