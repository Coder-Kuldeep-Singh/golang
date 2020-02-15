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
	Latitude   float64
	Longitude  float64
	population int
}

func getPopulations() {
	db := getDatabaseConnection()
	flags := flag.String("l", "0", "provide the limit of the data you want from database")
	flag.Parse()
	dbrows, err := db.Query("select latitude,longitude, population from offer order by population desc limit " + *flags)
	if err != nil {
		fmt.Println("BAD SQL", err)
	}
	// offers := []Offer{}
	for dbrows.Next() {
		var latitude float64
		var longitude float64
		var poplations int
		err = dbrows.Scan(&latitude, &longitude, &poplations)
		if err != nil {
			fmt.Println(err)
		}
		var population Populations
		population.Latitude = latitude
		population.Longitude = longitude
		population.population = poplations
		fmt.Println("Latitude : ", population.Longitude)
		fmt.Println("Longitude : ", population.Longitude)
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
