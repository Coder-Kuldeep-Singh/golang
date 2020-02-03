package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Homepage)
	fmt.Println("Development Server Started 127.0.0.1:9000")
	err := http.ListenAndServe(":9000", nil) //setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func Homepage(w http.ResponseWriter, req *http.Request) {
	// if req.URL.Path != "/" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }
	if req.Method == "GET" {
		http.ServeFile(w, req, "/templates/header.html")
	} else {
		fmt.Fprint(w, "Not Found")
	}

}
