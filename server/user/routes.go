package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/sample-tutorials/server/util"

	"github.com/gorilla/mux"
)

//RequireToken enables required Token verification on device
func (u *User) RequireToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if verifyToken(r.Header.Get("x-access-token"), r.Header.Get("x-access-userid")) {
		log.Println("called")
		next(w, r)
	} else {
		b, _ := json.Marshal(AddOutput{false, "Must login to perform this action", ""})
		w.Write(b)
	}
}

//RequireAdmin routes after calling this will require an admin account
func (u *User) RequireAdmin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if verifyToken(r.Header.Get("x-access-token"), r.Header.Get("x-access-userid")) {
		if u.isAdmin(r.Header.Get("x-access-userid")) {
			next(w, r)
		} else {
			b, _ := json.Marshal(AddOutput{false, "Must be admin to perform this action", ""})
			w.Write(b)
		}
	} else {
		b, _ := json.Marshal(AddOutput{false, "Must be logged in to perform this action", ""})
		w.Write(b)
	}
}

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
