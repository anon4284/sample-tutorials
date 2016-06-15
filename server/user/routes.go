package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/shir-website/server/util"

	"github.com/gorilla/mux"
)

//EnableSignup enables you to signup
func (u *User) EnableSignup(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		req := AddInput{}
		util.UnmarshalResponse(r.Body, &req)
		valid := signupValid(&req)
		if valid.Valid {
			if u.EmailUnique(req.Useremail) {
				o := u.Add(&req)
				b, _ := json.Marshal(o)
				w.Write(b)
			} else {
				b, _ := json.Marshal(AddOutput{false, "Email already registerd", ""})
				w.Write(b)
			}
		} else {
			b, _ := json.Marshal(valid)
			w.Write(b)
		}
	})
}

//EnableLogin enables login to application
func (u *User) EnableLogin(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		resp := LoginInput{}
		util.UnmarshalResponse(r.Body, &resp)
		login := u.Login(&resp)
		b, _ := json.Marshal(login)
		w.Write(b)
	})
}

//EnableChangePassword enables route to change password
func (u *User) EnableChangePassword(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
	})
}
