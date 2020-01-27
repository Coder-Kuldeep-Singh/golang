package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
       _"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //
	//fmt.Println("Method: ", r.Method) //get method of request
	//fmt.Println(r.Form) // print information on server side.
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme",r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:",k)
	//	fmt.Println("val:",strings.Join(v,""))
	//}
	fmt.Fprintf(w,"Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w,nil)
	} else {
		r.ParseForm()
		//logic part of login 
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/",sayhelloName) //setting router rules
	http.HandleFunc("/login",login)
	fmt.Println("Development Server Started 127.0.0.1:9000/login")
	err := http.ListenAndServe(":9000",nil) //setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}













