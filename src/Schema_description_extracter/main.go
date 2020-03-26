package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type JSONSCHEMA struct {
	Context string `json:"@context"`
	Graphs  []Subchilds
}
type Subchilds struct {
	Description string `json:"description"`
}

func fetchSchema(URL string, wg *sync.WaitGroup) {
	defer wg.Done()
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(URL)
	if err != nil {
		log.Println("having problem to find url", err)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		return
	}
	// log.Println(URL)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	output := ""
	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("script[type=\"application/ld+json\"]").Each(func(index int, item *goquery.Selection) {
		contents := item.Text()
		output = contents
	})
	if output == "" {
		log.Println("WARNING: No Schema FOUND on " + URL)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		return
	}
	Slice := []byte(output)
	var s Subchilds
	err = json.Unmarshal(Slice, &s)
	if err != nil {
		log.Println("Error while Unmarshling JSON data")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}
	fmt.Println(s.Description)
	// fmt.Println(output)
	fmt.Println("#######################################################################################################################################################")
}

func ReadFile(filepath string, wg *sync.WaitGroup) {
	readfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("Error to reading the file", err)
		os.Exit(1)
	}
	split := strings.Split(string(readfile), "\n")
	// wg.Add(len(split))
	for _, line := range split {
		// fmt.Println(line)
		wg.Add(1)
		go fetchSchema(line, wg)
	}
	// wg.Wait()
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	filepath := flag.String("f", "", "Provide the File name")
	url := flag.String("u", "", "Provide the url")
	flag.Parse()
	wg.Add(1)
	if *url != "" {
		go fetchSchema(*url, &wg)
	}
	if *filepath != "" {
		go ReadFile(*filepath, &wg)
	}
	wg.Wait()

}
