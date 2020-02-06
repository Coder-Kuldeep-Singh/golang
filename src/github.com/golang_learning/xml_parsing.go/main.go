package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func main() {
	domain := flag.String("d", "domain.com", "provide domain")
	protocol := flag.String("p", "https", "Protocol")
	flag.Parse()
	response, err := http.Get(*protocol + "://" + *domain + "/robots.txt")
	Error(err)
	defer response.Body.Close()
	responsebody, err := ioutil.ReadAll(response.Body)
	Error(err)
	str := string(responsebody)
	xmlfinder(str)
	// fmt.Println(str)
}

// func parseXmldata(xmldata string) {

// }
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func xmldatacollect(xmlfile string) {
	response, err := http.Get(xmlfile)
	Error(err)
	defer response.Body.Close()
	responsebody, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(responsebody))
	// parseXmldata(responsebody)
	var s SitemapIndex
	xml.Unmarshal(responsebody, &s)
	fmt.Println(s.Locations)
}

func xmlfinder(domain string) {
	sitemap := regexp.MustCompile(`sitemap: (.*)`)
	xmlpath := sitemap.FindAllStringSubmatch(domain, -1)
	// fmt.Println(xmlpath)
	for _, element := range xmlpath {
		// fmt.Println(element[1])
		xmldatacollect(element[1])
	}
}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
