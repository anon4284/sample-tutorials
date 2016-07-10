package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/sample-tutorials/server/util"
	"strconv"

	"github.com/gorilla/mux"
)

//RequireToken enables required Token verification on device
func (u *User) RequireToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if verifyToken(r.Header.Get("x-access-token"), r.Header.Get("x-access-userid")) {
		log.Println("called")
		next(w, r)
	} else {
		b, _ := json.Marshal(Output{false, "Must login to perform this action", ""})
		w.Write(b)
	}
}

//RequireAdmin routes after calling this will require an admin account
func (u *User) RequireAdmin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if verifyToken(r.Header.Get("x-access-token"), r.Header.Get("x-access-userid")) {
		if u.isAdmin(r.Header.Get("x-access-userid")) {
			next(w, r)
		} else {
			b, _ := json.Marshal(Output{false, "Must be admin to perform this action", ""})
			w.Write(b)
		}
	} else {
		b, _ := json.Marshal(Output{false, "Must be logged in to perform this action", ""})
		w.Write(b)
	}
}

func (u *User) EnableDelete(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		req := DeleteInput{}
		util.UnmarshalResponse(r.Body, &req)
		resp := u.Delete(req.UserID)
		b, _ := json.Marshal(resp)
		w.Write(b)
	})
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
				b, _ := json.Marshal(Output{false, "Email already registerd", ""})
				w.Write(b)
			}
		} else {
			b, _ := json.Marshal(valid)
			w.Write(b)
		}
	})
}

//EnableGetMe get user info
func (u *User) EnableGetMe(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		user := u.Get(r.Header.Get("x-access-userid"))
		me := MeInfo{*user.Item["useremail"].S, *user.Item["username"].S, *user.Item["profilePicture"].S}
		b, _ := json.Marshal(me)
		w.Write(b)
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

//EnableGetAll gets all users
func (u *User) EnableGetAll(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		myLimit, _ := strconv.Atoi(limit)
		users := u.GetAll(int64(myLimit))
		b, _ := json.Marshal(users)
		w.Write(b)
	})
}

//EnableChangePassword enables route to change password
func (u *User) EnableChangePassword(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		resp := ChangePasswordInput{}
		util.UnmarshalResponse(r.Body, &resp)
		pw := u.changePasswordValid(resp, r.Header.Get("x-access-userid"))
		b, _ := json.Marshal(pw)
		w.Write(b)
	})
}
