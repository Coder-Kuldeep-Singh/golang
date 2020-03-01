package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func fetch(urlpath string, wg *sync.WaitGroup) {
	response, err := http.Get(urlpath)
	if err != nil {
		log.Printf("%s", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Println("Host not found")
		log.Println(urlpath)
		return
	}
	//Parsing url
	domain, err := url.Parse(urlpath)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s", err)
		return
	}
	if string(body) == "" {
		log.Println("No Sitemap Found")
		log.Println(urlpath)
	}

	var s SitemapIndex
	xml.Unmarshal(body, &s)
	for _, Location := range s.Locations {
		wg.Add(1)
		fmt.Printf("%s\n", Location)
	}
	createfile(string(domain.Host), string(body), wg)
	wg.Wait()
	// IfBodyDataIsSimple(string(body))
}
func IfBodyDataIsSimple(pageContent string) {
	split := strings.Split(pageContent, "\n")
	log.Println(split)
	return
}

func createfile(filename, output string, wg *sync.WaitGroup) {
	defer wg.Done()
	out, err := os.Create("./output/" + filename + ".xml")
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

func ReadFile(filepath string, wg *sync.WaitGroup) {
	readfile, err := ioutil.ReadFile(filepath)
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
		wg.Add(1)
		go fetch(line, wg)
	}
	wg.Wait()
	// wg.Done()
}

func main() {
	var wg sync.WaitGroup
	filepath := flag.String("f", "", "Provide the path of the file")
	flag.Parse()
	wg.Add(1)
	go ReadFile(*filepath, &wg)
	wg.Wait()
	// fetch("https://chaufferus.us/https-sitemap_index.xml")

}
