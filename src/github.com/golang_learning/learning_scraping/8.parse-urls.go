package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

func main() {
	uri := flag.String("u", "", "Provide the url")
	flag.Parse()
	//make an http request
	response, err := url.Parse(*uri)
	if err != nil {
		log.Fatal(err)
	}
	//Print out URL pieces
	fmt.Println("Scheme: " + response.Scheme)
	fmt.Println("Host: " + response.Host)
	fmt.Println("Path: " + response.Path)
	fmt.Println("Query string: " + response.RawQuery)
	fmt.Println("Fragment: " + response.Fragment)

	//Get the query key|values as a map
	fmt.Println("\nQuery values:")
	querymap := response.Query()
	fmt.Println(querymap)

	//Create a new URL from scratch
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("key1", "value1")
	newQueryValues.Set("key2", "value2")
	customURL.Fragment = "bookmarkList"
	customURL.RawQuery = newQueryValues.Encode()
	fmt.Println("\nCustom URL:")
	fmt.Println(customURL.String())

}
