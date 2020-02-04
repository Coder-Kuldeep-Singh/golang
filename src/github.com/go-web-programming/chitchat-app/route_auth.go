package main

import (
	"fmt"
	"net/http"
)

func authenticate(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user, err := data.UserByEmail(req.PostFormValue("email"))
	Error(err)
	if user.Password == data.Encrypt(req.PostFormValur("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, req, "/", 302)
	} else {
		http.Redirect(w, req, "/login", 302)
	}
}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
