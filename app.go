package main

import (
	"projects/sample-tutorials/server/db"
	"projects/sample-tutorials/server/portfolio"
	"projects/sample-tutorials/server/routes"
	"projects/sample-tutorials/server/user"
)

func main() {

	myDB := db.New(db.NewConfigFromJSON("creds.json"))
	myUser := user.New(&myDB.SVC, "cg-users")
	myPortfolio := portfolio.New(&myDB.SVC, "st-portfolio")

	router := router.New(5000)
	//	routerVerified := mux.NewRouter()
	router.ServeHTMLifNotFound("./public/web/index.html")
	myUser.EnableLogin(router.Router, "/api/user/login")

	myPortfolio.EnableRouteAdd(router.Router, "/api/projects/add")
	myPortfolio.EnableRouteGetAll(router.Router, "/api/projects/get")
	myPortfolio.EnableRouteGetByProjectID(router.Router, "/api/projects/getByID")

	router.Start()
}
