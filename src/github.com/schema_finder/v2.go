package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func fetchJob(jobURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(jobURL)
	doc, err := goquery.NewDocument(jobURL)
	if err != nil {
		log.Println(err)
		return
	}
	//Parsing url
	dom, err := url.Parse(jobURL)
	if err != nil {
		log.Println(err)
		return
	}
	filename := dom.Host
	output := ""
	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("script[type=\"application/ld+json\"]").Each(func(index int, item *goquery.Selection) {
		contents := item.Text()
		output = contents
	})
	if output == "" {
		log.Println("WARNING: No Schema FOUND on " + jobURL)
	}
	// fmt.Println(output)
	out, err := os.Create("./output/" + filename + ".json")
	if err != nil {
		// return err
		log.Println(err)
		return
	}
	defer out.Close()
	//Write the body into file
	_, err = io.WriteString(out, output)
	if err != nil {
		// return err
		log.Println(err)
	}
}
func main() {
	var wg sync.WaitGroup
	path := flag.String("f", "", "Provide the File name")
	flag.Parse()
	readfile, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Println("Error to read file")
	}
	split := strings.Split(string(readfile), "\n")
	for _, line := range split {
		wg.Add(1)
		go fetchJob(line, &wg)
	}
	//No newline at end of file
	wg.Wait()
}
