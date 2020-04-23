package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Sitemaps struct {
	ID           int
	Sitemap_Urls string
	Domain       string
}

func InitDB() (*sql.DB, error) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetSitemapUrl() []Sitemaps {
	var output []Sitemaps
	db, err := InitDB()
	if err != nil {
		log.Println("Connection string failed! ", err)
		return output
	}
	query, err := db.Query(`SELECT id,sitemap_url,domain from robot`)
	if err != nil {
		log.Println("Query Failed!")
		return output
	}
	defer db.Close()
	for query.Next() {
		var sitemap, domain string
		var id int
		err = query.Scan(&id, &sitemap, &domain)
		if err != nil {
			log.Println("Error to Execute Statements")
			return output
		}
		var site = Sitemaps{id, sitemap, domain}
		output = append(output, site)
	}
	return output
}
