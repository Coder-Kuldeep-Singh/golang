package main

import (
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, req *http.Request) (sess data.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
