package main

// takes a json and returns a record.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ZipRecruiterAPIResult Represents the single record result from the ZipRecruiter API
type Resources struct {
	SalaryInterval     string  `json:"salary_interval"`
	Location           string  `json:"location"`
	Country            string  `json:"country"`
	PostedTimeFriendly string  `json:"posted_time_friendly"`
	SalaryMin          float64 `json:"salary_min"`
	Snippet            string  `json:"snippet"`
	SalaryMaxAnnual    float64 `json:"salary_max_annual"`
	ID                 string  `json:"id"`
	IndustryName       string  `json:"industry_name"`
	BuyerType          string  `json:"buyer_type"`
	LastPlanName       string  `json:"last_plan_name"`
	State              string  `json:"state"`
	PostedTime         string  `json:"posted_time"`
	JobAge             float64 `json:"job_age"`
	SalaryMax          float64 `json:"salary_max"`
	HiringCompany      struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ID          string `json:"id"`
		URL         string `json:"url"`
	} `json:"hiring_company"`
	URL             string  `json:"url"`
	SalaryMinAnnual float64 `json:"salary_min_annual"`
	SalarySource    string  `json:"salary_source"`
	Category        string  `json:"category"`
	HasNonZrURL     string  `json:"has_non_zr_url"`
	HasZipapply     bool    `json:"has_zipapply"`
	Name            string  `json:"name"`
	Source          string  `json:"source"`
	City            string  `json:"city"`
}

func main() {
	var jobs []Resources
	if len(os.Args) < 2 {
		log.Println("usage: " + os.Args[0] + "  <json-file>")
		return
	}
	fileName := os.Args[1]
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("File was not readable: " + err.Error())
		return
	}
	err = json.Unmarshal(bytes, &jobs)
	if err != nil {
		log.Println("Invalid set of results:" + err.Error())
		return
	}
	for j := range jobs {
		job := jobs[j]
		log.Println(job) // job title
		InsertAllJsonData(job)

	}

}

func InsertAllJsonData(schema Resources) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	mysqlConnectionstring := fmt.Sprintf(dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True")
	db, err := gorm.Open("mysql", mysqlConnectionstring)
	if err != nil {
		log.Println("Connection string is not valid", err)
	}
	defer db.Close()
	// db := databaseConnectionString()
	// fmt.Println(schema)

	// Check the table exists or not
	db.HasTable(schema)

	//if table exists delete table
	db.DropTable(schema)

	//create table
	db.CreateTable(schema)

	db.Model(schema).ModifyColumn("url", "text")

	//GORM insert statement here
	db.Debug().Create(schema)

}
