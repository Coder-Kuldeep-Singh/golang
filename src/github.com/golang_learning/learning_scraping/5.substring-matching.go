package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ImagesUrl struct {
	images string
}

var ImgMap = map[int]ImagesUrl{}

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
	Openingtag := "<style>"
	ClosingTag := "</style>"
	Length := len(Openingtag)
	StartIndex := strings.Index(pageContent, Openingtag)
	if StartIndex == -1 {
		fmt.Println("No element found")
		os.Exit(0)
	}
	//Find the index of the closing tag
	EndIndex := strings.Index(pageContent, ClosingTag)
	if EndIndex == -1 {
		fmt.Println("No closing tag found.")
		os.Exit(0)
	}

	pageImages := []byte(pageContent[StartIndex+Length : EndIndex])
	re := regexp.MustCompile(`background-image:(.*);`)
	bgimages := re.FindAllString(string(pageImages), -1)
	if bgimages == nil {
		fmt.Println("No matches.")
	} else {
		for i, bgimage := range bgimages {
			ImgMap[i] = ImagesUrl{images: bgimage}
			// ImgMap[i] = ImagesUrl{images: bgimage}
			// fmt.Println(bgimage)
			fmt.Printf("1 -> %v\n", ImgMap)
		}
	}
	// fmt.Printf("Page title: %s\n", pageTitle)
}
