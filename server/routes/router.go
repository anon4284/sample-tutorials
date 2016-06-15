package router

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//Router the main router of the application
type Router struct {
	Port   int
	Server *negroni.Negroni
	Router *mux.Router
}

//New create new Router
func New(port int) *Router {
	return &Router{Server: negroni.Classic(), Router: mux.NewRouter(), Port: port}
}

//ServeHTMLifNotFound serves html file if no other route was specified by the frontend
func (r *Router) ServeHTMLifNotFound(path string) {
	r.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
}

//EnableTestRoute enables route to test
func (r *Router) EnableTestRoute() {
	r.Router.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}

//AddRouter adds another router to the server
func (r *Router) AddRouter(router *mux.Router) {
	r.Server.UseHandler(router)
}

//Start assignes handler to Server and starts the server
func (r *Router) Start() {
	r.Server.UseHandler(r.Router)
	r.Server.Run(":" + strconv.Itoa(r.Port))
}
