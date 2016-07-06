package main

import (
	"projects/sample-tutorials/server/comments"
	"projects/sample-tutorials/server/db"
	"projects/sample-tutorials/server/portfolio"
	"projects/sample-tutorials/server/routes"
	"projects/sample-tutorials/server/user"

	"github.com/gorilla/mux"
)

func main() {

	myDB := db.New(db.NewConfigFromJSON("creds.json"))
	myUser := user.New(&myDB.SVC, "cg-users")
	myPortfolio := portfolio.New(&myDB.SVC, "st-portfolio")
	myComments := comments.New(&myDB.SVC, "st-comments", *myPortfolio, *myUser)

	router := router.New(5000)

	router.ServeHTMLifNotFound("./public/web/index.html")
	myUser.EnableLogin(router.Router, "/api/user/login")
	myUser.EnableSignup(router.Router, "/api/user/signup")
	myPortfolio.EnableRouteGetAll(router.Router, "/api/projects/get")
	myPortfolio.EnableRouteGetByProjectID(router.Router, "/api/projects/getByID")

	myComments.EnableRouteGetByProjectID(router.Router, "/api/comments/getByProjectID")

	userRouter := mux.NewRouter()
	router.AddRouter(userRouter, "/api/user", myUser.RequireToken)

	myComments.EnableRouteAdd(userRouter, "/api/user/comments/add")

	adminRouter := mux.NewRouter()
	router.AddRouter(adminRouter, "/api/admin", myUser.RequireAdmin)

	myPortfolio.EnableRouteAdd(adminRouter, "/api/admin/projects/add")

	router.Start()
}
