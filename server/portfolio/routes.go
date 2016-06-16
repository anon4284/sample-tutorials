package portfolio

import (
	"encoding/json"
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
