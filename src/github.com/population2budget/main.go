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
type Populations struct {
	ID         int
	Date       string
	Jobtitle   string
	City       string
	population int
}

func getPopulations() {
	db := getDatabaseConnection()
	dbrows, err := db.Query("select id,date, jobtitle, city, population from Populations")
	if err != nil {
		fmt.Println("BAD SQL", err)
	}
	// offers := []Offer{}
	for dbrows.Next() {
		var id int
		var date string
		var jobTitle string
		var city string
		var poplations int
		err = dbrows.Scan(&id, &date, &jobTitle, &city, &poplations)
		if err != nil {
			fmt.Println(err)
		}
		var population Populations
		population.ID = id
		population.Date = date
		population.Jobtitle = jobTitle
		population.City = city
		population.population = poplations
		// fmt.Println(population.ID, "------->", population.Date, "------->", population.Jobtitle, "------->", population.City, "------->", population.population)
		fmt.Println("ID : ", population.ID)
		fmt.Println("Date : ", population.Date)
		fmt.Println("JobTitle : ", population.Jobtitle)
		fmt.Println("City : ", population.City)
		fmt.Println("Population : ", population.population)
		fmt.Println("***************************************************************************")
		// offers = append(offers, offer)
	}
	// fmt.Println(offers)
	db.Close()
}
func main() {
	getPopulations()
}
