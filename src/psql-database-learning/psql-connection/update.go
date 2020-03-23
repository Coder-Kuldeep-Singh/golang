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
	sqlStatement = `
		UPDATE users
		SET first_name = $2, last_name = $3
		WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 6, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
