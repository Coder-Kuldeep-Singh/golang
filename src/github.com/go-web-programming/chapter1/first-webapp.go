package main
import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage %s", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/",homepage)
	fmt.Println("Server Running 127.0.0.1:8000")
	http.ListenAndServe(":8000",nil)
}

