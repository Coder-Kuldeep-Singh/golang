package main

import (
	"database/sql"
	"fmt"
	//	"io"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

)

func getDatabaseConnection() *sql.DB {
	connectionString := "root:@tcp(localhost:3306)/deliveryjobsnyc"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
func checkErr(e error) {
	if e == nil {
		panic(e)
	}
}

// Offer represents a job
type Offer struct {
	ID             int
	JobTitle       string
	JobDescription string
}

func getOffers() ([]Offer, error) {
	db := getDatabaseConnection()
	dbrows, err := db.Query("select id, job_title, job_description from site")
	if err != nil {
		fmt.Println("BAD SQL", err)
		return nil, err
		// io.WriteString(w, "Not Found records")
		// return []
		//return nil, err
	}
	offers := []Offer{}
	for dbrows.Next() {
		var id int
		var jobTitle string
		var jobDescription string
		err = dbrows.Scan(&id, &jobTitle, &jobDescription)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		var offer Offer
		offer.ID = id
		offer.JobTitle = jobTitle
		offer.JobDescription = jobDescription
		offers = append(offers, offer)
	}
	db.Close()
	// var m map[string]Vertex
	return offers, nil
}
func main() {
	offers, err := getOffers()
	if err != nil {
		log.Println("Failed:" + err.Error())
		return
	}
	header, err := ioutil.ReadFile("./templates/header.html")
	if err != nil {
		log.Println("header Failed")
		return
	}

	footer, err := ioutil.ReadFile("./templates/footer.html")
	if err != nil {
		log.Println("footer Failed")
		return
	}

	for offerID := range offers {
		offer := offers[offerID]
		log.Println(offer.ID, offer.JobTitle)
		d1 := []byte{}
		d1 = append(d1, header...)
		d1 = append(d1, []byte(offer.JobDescription)...)
		d1 = append(d1, footer...)
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, "(", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, ")", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, " ", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, "*", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, "?", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, ":", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, ";", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, "/", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, "\\", "-")
		offer.JobTitle = strings.ReplaceAll(offer.JobTitle, ",", "-")
		fileName := fmt.Sprintf("./output/%d--%s.html", offer.ID, offer.JobTitle)

		err := ioutil.WriteFile(fileName, d1, 0755)
		if err != nil {
			log.Println("Failed to write on disk" + err.Error())

		}
	}
}
