package main

import (
	"database/sql"
	"flag"
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

// Population
type Populations struct {
	ID int
	// Date       string
	// Jobtitle   string
	// City       string
	population int
	Zipcode    int
}

func getPopulations() {
	db := getDatabaseConnection()
	flags := flag.String("l", "0", "provide the limit of the data you want from database")
	flag.Parse()
	dbrows, err := db.Query("select id, population,zipcode from offer order by population desc limit " + *flags)
	if err != nil {
		fmt.Println("BAD SQL", err)
	}
	// offers := []Offer{}
	for dbrows.Next() {
		var id int
		// var date string
		// var jobTitle string
		// var city string
		var poplations int
		var zipcode int
		err = dbrows.Scan(&id, &poplations, &zipcode)
		if err != nil {
			fmt.Println(err)
		}
		var population Populations
		population.ID = id
		// population.Date = date
		// population.Jobtitle = jobTitle
		// population.City = city
		population.population = poplations
		population.Zipcode = zipcode
		// fmt.Println(population.ID, "------->", population.Date, "------->", population.Jobtitle, "------->", population.City, "------->", population.population)
		fmt.Println("ID : ", population.ID)
		// fmt.Println("Date : ", population.Date)
		// fmt.Println("JobTitle : ", population.Jobtitle)
		// fmt.Println("City : ", population.City)
		fmt.Println("Zipcode : ", population.Zipcode)
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
