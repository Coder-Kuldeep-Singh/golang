package main

import (
	//"bufio"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open the file
	filename := flag.String("f", "", "provide path of the file")
	flag.Parse()
	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	db := getDatabaseConnection()
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		sqlStatement := `INSERT INTO Populations (date, jobtitle, city, population)
							VALUES (?, ?, ?, ?)`
		_, err = db.Exec(sqlStatement, record[0], record[1], record[2], record[3])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data Inserted")
	os.Exit(0)
}

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
