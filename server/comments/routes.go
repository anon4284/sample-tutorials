package comments

import (
	"encoding/json"
	"net/http"
	"projects/babylongo/server/util"

	"github.com/gorilla/mux"
)

//EnableRouteAdd route to add portfolio items
func (c *Comments) EnableRouteAdd(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		req := AddInput{}
		util.UnmarshalResponse(r.Body, &req)
		o := c.Add(&req, r.Header.Get("x-access-userid"))
		b, _ := json.Marshal(o)
		w.Write(b)
	})
}

//EnableRouteGetByProjectID gets all comments of project
func (c *Comments) EnableRouteGetByProjectID(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		o := c.GetByProjectID(r.URL.Query().Get("id"))
		b, _ := json.Marshal(o)
		w.Write(b)
	})
}
