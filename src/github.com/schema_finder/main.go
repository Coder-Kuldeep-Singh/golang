package main

import (
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ReadFile(path string) {
	readfile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error to reading the file", err)
		os.Exit(1)
	}
	split := strings.Split(string(readfile), "\n")
	for _, line := range split {
		// timeout := time.Duration(1 * time.Second)
		// conn, err := net.DialTimeout("tcp","mysyte:myport", timeout)
		// if err != nil {
		// 		log.Println("Site unreachable, error: ", err)
		// }
		parsingUrlTocheck(line)
	}
}
func parsingUrlTocheck(urll string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(urll)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	// if res.StatusCode != 200 {
	//   log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Problem to find url", err)
		return
	}
	// fmt.Println(doc)
	//Parsing url
	dom, err := url.Parse(urll)
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
		log.Println("WARNING: No Schema FOUND!!!")
	}
	createfile(filename, output)

}

func main() {
	path := flag.String("f", "", "Provide the File name")
	flag.Parse()
	ReadFile(*path)
	// fmt.Println(output)
}

func createfile(filename, output string) {
	out, err := os.Create(filename + ".json")
	if err != nil {
		log.Println(err)
		return

	}
	defer out.Close()
	//Write the body into file
	_, err = io.WriteString(out, output)
	if err != nil {
		log.Println(err)
		return
	}
}
