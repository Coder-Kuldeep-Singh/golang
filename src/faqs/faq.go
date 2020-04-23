package main

import (
	"html/template"
	"net/http"
)

func HtmlOutput(w http.ResponseWriter, r *http.Request) {
	data := GetSitemapUrl()
	t, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	http.HandleFunc("/", HtmlOutput)
	http.ListenAndServe(":8080", nil)
}
