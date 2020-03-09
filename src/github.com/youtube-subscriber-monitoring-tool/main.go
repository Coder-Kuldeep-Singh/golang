package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/youtube-subscriber-monitoring-tool/youtube-stats/websocket"
)

// homepage will be a simple "hello world" style page
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// our new stats function which will expose any YouTube
// stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {
	// we call our new websocket package Upgrade
	// function in order to upgrade the connection
	// from a standard HTTP connection to a websocket one
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	// we then call our Writer function
	// which continually polls and writes the results
	// to this websocket connection
	go websocket.Writer(ws)
}

// setup handles setting up our servers
// routes and matching them to their respective
// functions
func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/stats", stats)
	// here we kick off our server on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// our main function
func main() {
	fmt.Println("Youtube Subscriber Monitor..")

	// calls setup routes
	setupRoutes()
}
