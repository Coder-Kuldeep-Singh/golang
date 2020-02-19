package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func processElement(index int, element *goquery.Selection) {
	//see if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func main() {
	url := flag.String("u", "", "Provide the url")
	flag.Parse()
	//make an http request
	response, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//read response data in memory
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body", err.Error())
		os.Exit(1)
	}
	document.Find("a").Each(processElement)

}
