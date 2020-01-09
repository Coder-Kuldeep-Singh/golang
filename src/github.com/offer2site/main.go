package main

import (
	"database/sql"
	_ "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Offer struct {
	offerid int
}

type Site struct {
	siteid int
}

func databaseconnectionvalues() {

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1)/deliveryjobsnyc")
	if err != nil {
		log.Print("Error in connection")
		log.Fatal(err)
	}
	log.Print("connected")
	defer db.Close()

	var (
		offer_id int
		site_id  int
	)
	rows, err := db.Query("select offer.id, site.id from offer,site")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&offer_id, &site_id)

		if err != nil {
			log.Print("Error to scan")
			log.Fatal(err)
		}
		log.Print("**********************************************************************************************************************************************************")
		log.Print("<tr>")

		eachValues := Offer{
			offerid: site_id}
		sitevalues := Site{
			siteid: offer_id}

		log.Println("<td>", "id : Site ", sitevalues, "</td>")
		log.Println("<td>", "id : Offer ", eachValues, "</td>")

		log.Print("</tr>")
		// insert data into database
		// stmt, err := db.Prepare("INSERT INTO offer2site(offer_id,site_id) VALUES(?,?)")
		// if err != nil {
		// 	log.Fatal(err)
		// 	log.Println("Error to Execute Query")
		// }
		// res, err := stmt.Exec(sitevalues, eachValues)
		// if err != nil {
		// 	log.Fatal(err)
		// 	log.Println("Error to execute Insert Query")
		// }
		// log.Println("Values Inserted Successfully ", res)
		// lastId, err := res.LastInsertId()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// rowCnt, err := res.RowsAffected()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	}
	err = rows.Err()
	if err != nil {
		log.Print("Error to find data")
		log.Fatal(err)
	}

}

func main() {
	databaseconnectionvalues()
}
