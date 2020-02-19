package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//parsing some flags for intractivity with CLI
	url := flag.String("u", "", "Provide the Url")
	filename := flag.String("f", "", "Provide the name of file to store data")
	flag.Parse()

	//make an http request to url|domain
	response, err := http.Get(*url)
	if err != nil {
		fmt.Println("Url Doesn't Exists", err)
	}
	defer response.Body.Close()

	//create the file to store data of url
	outfile, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	//Copy data from http response to file
	_, err = io.Copy(outfile, response.Body)
	if err != nil {
		fmt.Println("Response Not found from the page")
	}
	fmt.Println("Successfully data inserted into file")
}
