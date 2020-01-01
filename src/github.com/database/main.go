package main

import (
	"database/sql"
	_ "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func database_connection_values() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Print("Error in connection")
		log.Fatal(err)
	}
	log.Print("connected")
	defer db.Close()
	var (
		id                      int
		name                    string
		site                    string
		site_url                string
		company_name            string
		job_title               string
		job_description         string
		educationalRequirements string
		experianceRequirements  string
		qualification           string
		responsibilities        string
		skills                  string
		value_hour              int
		sid                     int
		folder                  string
		offer_module            string
		enabled                 int
		destination             string
		organization            string
		occupational_category   string
		organization_logo       string
		script_template         string
	)
	// var (
	// 	id                      int
	// 	name                    string
	// 	site                    string
	// 	site_url                string
	// 	company_name            string
	// 	job_title               string
	// 	job_description         string
	// 	educationalRequirements string
	// 	experianceRequirements  string
	// 	qualification           string
	// 	responsibilities        string
	// 	skills                  string
	// 	value_hour              int
	// 	sid                     int
	// 	folder                  string
	// 	offer_module            string
	// 	enabled                 int
	// 	destination             string
	// 	organization            string
	// 	occupational_category   string
	// 	organization_logo       string
	// 	script_template         string
	// )
	// rows, err := db.Query("select chaufferjob.site.id,chaufferjob.site.name,chaufferjob.site.site,chaufferjob.site.site_url,chaufferjob.site.company_name,chaufferjob.site.job_title,chaufferjob.site.job_description,chaufferjob.site.educationRequirements,chaufferjob.site.experienceRequirements,chaufferjob.site.qualifications,chaufferjob.site.responsibilities,chaufferjob.site.skills,chaufferjob.site.value_hour,chaufferjob.site.sid,chaufferjob.site.folder,chaufferjob.site.offer_modulus,chaufferjob.site.enabled,chaufferjob.site.destination,chaufferjob.site.organization,chaufferjob.site.occupational_category,chaufferjob.site.organization_logo,chaufferjob.site.script_template , deliveryjobsnyc.site.id,deliveryjobsnyc.site.name,deliveryjobsnyc.site.site,deliveryjobsnyc.site.site_url,deliveryjobsnyc.site.company_name,deliveryjobsnyc.site.job_title,deliveryjobsnyc.site.job_description,deliveryjobsnyc.site.educationRequirements,deliveryjobsnyc.site.experienceRequirements,deliveryjobsnyc.site.qualifications,deliveryjobsnyc.site.responsibilities,deliveryjobsnyc.site.skills,deliveryjobsnyc.site.value_hour,deliveryjobsnyc.site.sid,deliveryjobsnyc.site.folder,deliveryjobsnyc.site.offer_modulus,deliveryjobsnyc.site.enabled,deliveryjobsnyc.site.destination,deliveryjobsnyc.site.organization,deliveryjobsnyc.site.occupational_category,deliveryjobsnyc.site.organization_logo,deliveryjobsnyc.site.script_template  from site")
	rows, err := db.Query("select chaufferjob.site.id,chaufferjob.site.name,chaufferjob.site.site,chaufferjob.site.site_url,						chaufferjob.site.company_name,chaufferjob.site.job_title,chaufferjob.site.job_description,								chaufferjob.site.educationRequirements,chaufferjob.site.experienceRequirements,											chaufferjob.site.qualifications,chaufferjob.site.responsibilities,chaufferjob.site.skills,								chaufferjob.site.value_hour,chaufferjob.site.sid,chaufferjob.site.folder,												chaufferjob.site.offer_modulus,chaufferjob.site.enabled,chaufferjob.site.destination,									chaufferjob.site.organization,chaufferjob.site.occupational_category,													chaufferjob.site.organization_logo,chaufferjob.site.script_template , deliveryjobsnyc.site.id,							deliveryjobsnyc.site.name,deliveryjobsnyc.site.site,deliveryjobsnyc.site.site_url,										deliveryjobsnyc.site.company_name,deliveryjobsnyc.site.job_title,														deliveryjobsnyc.site.job_description,deliveryjobsnyc.site.educationRequirements,										deliveryjobsnyc.site.experienceRequirements,deliveryjobsnyc.site.qualifications,										deliveryjobsnyc.site.responsibilities,deliveryjobsnyc.site.skills,														deliveryjobsnyc.site.value_hour,deliveryjobsnyc.site.sid,deliveryjobsnyc.site.folder,									deliveryjobsnyc.site.offer_modulus,deliveryjobsnyc.site.enabled,deliveryjobsnyc.site.destination,						deliveryjobsnyc.site.organization,deliveryjobsnyc.site.occupational_category,											deliveryjobsnyc.site.organization_logo,deliveryjobsnyc.site.script_template  from chaufferjob.site, deliveryjobsnyc.site")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &site, &site_url, &company_name, &job_title, &job_description, &educationalRequirements, &experianceRequirements, &qualification, &responsibilities, &skills, &value_hour, &sid, &folder, &offer_module, &enabled, &destination, &organization, &occupational_category, &organization_logo, &script_template, &id, &name, &site, &site_url, &company_name, &job_title, &job_description, &educationalRequirements, &experianceRequirements, &qualification, &responsibilities, &skills, &value_hour, &sid, &folder, &offer_module, &enabled, &destination, &organization, &occupational_category, &organization_logo, &script_template)
		if err != nil {
			log.Print("Error to scan")
			log.Fatal(err)
		}
		log.Print("**********************************************************************************************************************************************************")
		log.Print("<tr>")
		// log.Println(id, name, site, site_url, company_name, job_title, job_description, educationalRequirements, experianceRequirements, qualification, responsibilities, skills, value_hour, sid, folder, offer_module, enabled, destination, organization, occupational_category, organization_logo, script_template)
		log.Println("id :", id)
		log.Println("name :", name)
		log.Println("site :", site)
		log.Println("site_url :", site_url)
		log.Println("company_url :", company_name)
		log.Println("job_title :", job_title)
		log.Println("job_description :", job_description)
		log.Println("educationalRequirements :", educationalRequirements)
		log.Println("experianceRequirements :", experianceRequirements)
		log.Println("qualification :", qualification)
		log.Println("responsibilities :", responsibilities)
		log.Println("skills :", skills)
		log.Println("value_hour :", value_hour)
		log.Println("sid :", sid)
		log.Println("folder :", folder)
		log.Println("offer_module :", offer_module)
		log.Println("enabled :", enabled)
		log.Println("destination :", destination)
		log.Println("organization :", organization)
		log.Println("occupational_category :", occupational_category)
		log.Println("organization_logo :", organization_logo)
		log.Println("script_template :", script_template)

		log.Print("</tr>")

	}
	err = rows.Err()
	if err != nil {
		log.Print("Error to find data")
		log.Fatal(err)
	}
}

func main() {
	database_connection_values()
}
