package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getDatabaseConnection() *sql.DB {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")
	connectionString := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Connected with %s\n", dbname)
	return db
}

func CheckErrors(err error) {
	if err == nil {
		panic(err)
	}
}

// Offer represents a job
type Offer struct {
	ID             int
	JobTitle       string
	JobDescription string
}

func getOffers() {
	db := getDatabaseConnection()
	dbrows, err := db.Query("select id, job_title, job_description from site")
	if err != nil {
		fmt.Println("BAD SQL", err)
	}
	// offers := []Offer{}
	for dbrows.Next() {
		var id int
		var jobTitle string
		var jobDescription string
		err = dbrows.Scan(&id, &jobTitle, &jobDescription)
		if err != nil {
			fmt.Println(err)
		}
		var offer Offer
		offer.ID = id
		offer.JobTitle = jobTitle
		offer.JobDescription = jobDescription
		fmt.Println(offer.ID, "------->", offer.JobTitle, "------->", offer.JobDescription)
		// offers = append(offers, offer)
	}
	// fmt.Println(offers)
	db.Close()
}
func main() {
	getOffers()
}
