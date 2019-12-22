package main
import (
	"fmt"
	"log"
	"net/http"
	"crypto/subtle"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8010"
	ADMIN_USER = "admin"
	ADMIN_PASSWORD = "admin"
)


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Golang")
}

func BasicAuth(handler http.HandleFunc, realm string) http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request)
	{
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTImeCompare([]byte(user)),[]byte(ADMIN_USER)) != 1 ||subtle.ConstantTImeCompare([]byte(pass),[]byte(ADMIN_PASSWORD)) != 1
		{
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth( helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server :", err)
		return
	}
}
