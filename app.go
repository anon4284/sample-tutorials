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
	myComments := comments.New(&myDB.SVC, "st-comments2", *myPortfolio, *myUser)

	router := router.New(5000)

	router.ServeHTMLifNotFound("./public/build/web/index.html")
	myUser.EnableLogin(router.Router, "/api/login")
	myUser.EnableSignup(router.Router, "/api/signup")
	myPortfolio.EnableRouteGetAll(router.Router, "/api/projects/get")
	myPortfolio.EnableRouteGetByProjectID(router.Router, "/api/projects/getByID")

	myComments.EnableRouteGetByProjectID(router.Router, "/api/comments/getByProjectID")

	userRouter := mux.NewRouter()
	router.AddRouter(userRouter, "/api/user", myUser.RequireToken)

	myUser.EnableGetAll(userRouter, "/api/user/getAll")
	myUser.EnableGetMe(userRouter, "/api/user/me")
	myUser.EnableChangePassword(userRouter, "/api/user/password/change")

	myComments.EnableRouteAdd(userRouter, "/api/user/comments/add")
	myPortfolio.EnableRouteAdd(userRouter, "/api/user/projects/add")

	adminRouter := mux.NewRouter()
	router.AddRouter(adminRouter, "/api/admin", myUser.RequireAdmin)

	myUser.EnableDelete(adminRouter, "/api/admin/user/delete")

	router.Start()
}
