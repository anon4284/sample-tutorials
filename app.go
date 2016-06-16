package main

import (
	"projects/sample-tutorials/server/db"
	"projects/sample-tutorials/server/routes"
	"projects/sample-tutorials/server/user"
)

func main() {

	myDb := db.New(db.NewConfigFromJSON("creds.json"))
	myUser := user.New(&myDb.SVC, "cg-users")
	myUser.CreateTable()
	myUser.LimitAmount(1)

	router := router.New(5000)
	router.ServeHTMLifNotFound("./public/web/index.html")
	router.EnableTestRoute()

	myUser.EnableSignup(router.Router, "/api/user/signup")
	myUser.EnableLogin(router.Router, "/api/user/login")
	myUser.EnableChangePassword(router.Router, "/api/user/password")

	router.Start()
}
