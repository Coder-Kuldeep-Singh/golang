package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//parsing some flags for intractivity with CLI
	url := flag.String("u", "", "Provide the Url")
	flag.Parse()

	//Create http client with time request
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//make an http request to url|domain
	response, err := client.Get(*url)
	if err != nil {
		fmt.Println("Url Doesn't Exists", err)
	}
	defer response.Body.Close()

	bytes, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		fmt.Println("Response Not found from the page")
	}
	log.Println("Number of bytes copied to STDOUT:", bytes)
}
