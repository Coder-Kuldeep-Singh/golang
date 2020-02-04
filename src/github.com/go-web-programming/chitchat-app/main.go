package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, req)
		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	//All urls
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newthread)
	mux.HandleFunc("/thread/creat", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	//server
	server := &http.Server{
		Addr:    "196.168.30.27:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
