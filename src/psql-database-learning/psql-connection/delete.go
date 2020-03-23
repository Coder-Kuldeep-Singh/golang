package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tutree"
	password = "12qwaszx"
	dbname   = "test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Connection string is not valid", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;`
	Alert, err := db.Exec(sqlStatement, 8)
	if err != nil {
		log.Println("Delete Statement Have Problem", err)
		return
	}
	fmt.Println(Alert)
}
