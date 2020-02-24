package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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

	//read response data in memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body", err.Error())
		os.Exit(1)
	}
	//create regular expression to find html tags and contents
	// re := regexp.MustCompile("<!--(.|\n)*?-->")
	// re := regexp.MustCompile("<img (.*)*/>")
	re := regexp.MustCompile(`<style>
	(.*)
	</style>`)
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
}
