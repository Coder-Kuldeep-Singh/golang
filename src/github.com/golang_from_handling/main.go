package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "strings"

	_ "github.com/go-sql-driver/mysql"
)

func Indexpage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index page")

}

func database_connection_value() (db *sql.DB) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASSWORD")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp"+"("+dbhost+":"+dbport+")"+"/"+dbname)
	if err != nil {
		fmt.Println("Error in connection")
		log.Fatal(err)
	}
	// defer db.Close()
	fmt.Println("connected")
	return db
}
func login(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("method:", r.Method) //get request method
	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch req.Method {
	case "GET":
		http.ServeFile(w, req, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", req.PostForm)
		comment_text := req.FormValue("comment_text")
		organization := req.FormValue("organization")
		insertStatment(comment_text, organization)
		fmt.Fprintf(w, "Comment_text = %s\n", comment_text)
		http.Redirect(w, req, "/", 301)
		// fmt.Fprintf(w, "Password = %s\n", password)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func insertStatment(comment_text string, organization string) {
	db := database_connection_value()
	// dbtable := os.Getenv("DBTABLE")
	insForm, err := db.Prepare("INSERT INTO comment(comment_text, organization) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(comment_text, organization)
	log.Println("INSERT: Comment_text: " + comment_text)
	defer db.Close()
}

func main() {
	// http.HandleFunc("/", hello)
	// fmt.Printf("Starting server for testing HTTP POST...\n")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	//     log.Fatal(err)
	// }
	http.HandleFunc("/Indexpage", Indexpage) //setting router rules
	// http.HandleFunc("/ping", database_connection_value)
	http.HandleFunc("/", login)
	fmt.Println("Development Server Started 127.0.0.1:9000/")
	err := http.ListenAndServe(":9000", nil) //setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
