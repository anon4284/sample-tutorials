package portfolio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/sample-tutorials/server/util"

	"github.com/gorilla/mux"
)

//EnableRouteAdd route to add portfolio items
func (p *Portfolio) EnableRouteAdd(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		req := AddInput{}
		util.UnmarshalResponse(r.Body, &req)
		o := p.Add(&req, r.Header.Get("x-access-userid"))
		b, _ := json.Marshal(o)
		w.Write(b)
	})
}

//EnableRouteGetByUser gets items by userID
func (p *Portfolio) EnableRouteGetByUser(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		o := p.GetByUser(r.Header.Get("x-access-userid"))
		b, _ := json.Marshal(o.Items)
		w.Write(b)
	})
}

//EnableRouteGetAll gets all items
func (p *Portfolio) EnableRouteGetAll(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		o := p.Scan()
		b, _ := json.Marshal(o)
		w.Write(b)
	})
}

func (p *Portfolio) EnableRouteGetByProjectID(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query().Get("id"))
		o := p.GetByProjectID(r.URL.Query().Get("id"))
		b, _ := json.Marshal(o)
		w.Write(b)
	})
}
