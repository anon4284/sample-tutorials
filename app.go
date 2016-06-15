package main

import (
	"projects/gobook/server/db"
	"projects/shir-website/server/routes"
	"projects/shir-website/server/user"
)

func main() {

	myDb := db.New(db.NewConfigFromJSON("creds.json"))
	myUser := user.New(&myDb.SVC, "cg-users")
	myUser.CreateTable()
	myUser.LimitAmount(1)

	router := router.New(7016)
	router.ServeHTMLifNotFound("./public/web/index.html")
	router.EnableTestRoute()

	myUser.EnableSignup(router.Router, "/api/user/signup")
	myUser.EnableLogin(router.Router, "/api/user/login")
	myUser.EnableChangePassword(router.Router, "/api/user/password")

	router.Start()
}
