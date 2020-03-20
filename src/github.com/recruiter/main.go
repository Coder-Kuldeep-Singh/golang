package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	Country              sql.NullString
	Salary_max_annual    sql.NullInt64
	Has_non_zr_url       sql.NullInt64
	Hiring_company       sql.NullString
	Industry_name        sql.NullString
	Salary_max           sql.NullInt64
	ID                   string
	Name                 sql.NullString
	Category             sql.NullString
	Salary_source        sql.NullString
	Source               sql.NullString
	Salary_interval      sql.NullString
	State                sql.NullString
	Has_zipapply         sql.NullString
	Salary_min_annual    sql.NullInt64
	Posted_time_friendly sql.NullString
	Buyer_type           sql.NullString
	City                 sql.NullString
	Location             sql.NullString
	Posted_time          sql.NullString
	Job_age              sql.NullInt64
	Last_plan_name       sql.NullString
	Salary_min           sql.NullInt64
	Snippet              sql.NullString
	URL                  sql.NullString
	Description          sql.NullString
	Logo                 sql.NullString
}

func DatabaseConnectionString() (db *sql.DB) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname)
	if err != nil {
		log.Println("Connection String failed", err)
	}
	// fmt.Println("connected")
	return db
}

//SelectAllUrls function selecting all from Sitemap table
func SelectAllURLS(wg *sync.WaitGroup) {
	defer wg.Done()

	db := DatabaseConnectionString()

	//Select data from ziprecruiter Table
	results, err := db.Query(`SELECT country,salary_max_annual,has_non_zr_url,hiring_company,
	industry_name, salary_max,id,name,category,salary_source,source,salary_interval,state,has_zipapply, salary_min_annual,
	posted_time_friendly,buyer_type,city,location,posted_time,
	job_age,last_plan_name,salary_min,snippet,url,description,logo FROM ziprecruiter`)
	CheckError(err, "Bad SQL")
	for results.Next() {
		var tag Tag

		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.Country, &tag.Salary_max_annual, &tag.Has_non_zr_url, &tag.Hiring_company, &tag.Industry_name, &tag.Salary_max, &tag.ID, &tag.Name, &tag.Category, &tag.Salary_source, &tag.Source, &tag.Salary_interval,
			&tag.State, &tag.Has_zipapply, &tag.Salary_min_annual, &tag.Posted_time_friendly, &tag.Buyer_type, &tag.City, &tag.Location, &tag.Posted_time, &tag.Job_age, &tag.Last_plan_name, &tag.Salary_min, &tag.Snippet,
			&tag.URL, &tag.Description, &tag.Logo)
		CheckError(err, "Error while scanning url's to fetch from database")
		log.Println(tag)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	}
	defer db.Close()

}
func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
		return
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	SelectAllURLS(&wg)
	wg.Wait()

}
