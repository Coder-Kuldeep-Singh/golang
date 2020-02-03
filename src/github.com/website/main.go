package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", Homepage)
	fmt.Println("Development Server Started 127.0.0.1:9000")
	err := http.ListenAndServe(":9000", nil) //setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func Homepage(w http.ResponseWriter, req *http.Request) {
	titles := Page{Title: "Home Page"}
	homepage, err := template.ParseFiles("./templates/header.html")
	Error(err)
	homepage.Execute(w, titles)
}

// func Homepage(w http.ResponseWriter, req *http.Request) {
// 	titles := Page{Title: "Home Page"}
// 	homepage, err := template.ParseFiles("./templates/header.html")
// 	Error(err)
// 	homepage.Execute(w, titles)
// }

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
