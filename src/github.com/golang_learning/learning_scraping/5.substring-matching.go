package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := flag.String("u", "", "Provide the url")
	flag.Parse()
	//make an http request
	response, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	//find a substr
	titleStartIndex := strings.Index(pageContent, "<head>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	//Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</head>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	fmt.Printf("Page title: %s\n", pageTitle)
}
