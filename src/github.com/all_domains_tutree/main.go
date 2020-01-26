package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Domains(w http.ResponseWriter, req *http.Request) {
	//fetch data from given url
	response, err := http.Get("http://s.tutree.com:7635/v1/groups")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	// Read data from url
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	storeDomains := string(body)
	trimdata := strings.Split(storeDomains, "\n")
	for _, url := range trimdata {
		response, err := http.Get("http://s.tutree.com:7635/v1/" + url)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		// Read data from url
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		storeDomains := string(body)
		trimdata := strings.Split(storeDomains, "\n")
		for _, url := range trimdata {
			fmt.Fprintln(w, url)

		}
	}
}

func main() {
	http.HandleFunc("/v1/domains", Domains)
	fmt.Println("Development Server Started on localhost:8000/v1/domains")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
