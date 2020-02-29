package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	path := flag.String("f", "", "Provide the File name")
	flag.Parse()
	readfile, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Println("Error to reading the file", err)
	}
	split := strings.Split(string(readfile), "\n")
	for _, line := range split {
		doc, err := goquery.NewDocument(line)
		if err != nil {
			log.Fatal(err)
		}
		//Parsing url
		dom, err := url.Parse(line)
		if err != nil {
			log.Println(err)
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
			log.Println("WARNING: No Schema FOUND!!!")
		}
		// fmt.Println(output)
		out, err := os.Create(filename + ".txt")
		if err != nil {
			// return err
			log.Println(err)

		}
		defer out.Close()
		//Write the body into file
		_, err = io.WriteString(out, output)
		if err != nil {
			// return err
			log.Println(err)
		}
	}
}
